package actors

import (
	"log"

	"github.com/gorilla/websocket"
	"github.com/asynkron/protoactor-go/actor"
)

type WebSocketActor struct {
	conn     *websocket.Conn
	chatRoom *actor.PID
}

func NewWebSocketActor(conn *websocket.Conn, chatRoom *actor.PID) actor.Producer {
	return func() actor.Actor {
		return &WebSocketActor{
			conn:     conn,
			chatRoom: chatRoom,
		}
	}
}

func (wsa *WebSocketActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *actor.Started:
		log.Println("WebSocket actor started")
		context.Send(wsa.chatRoom, &Register{Client: context.Self()})
		go wsa.readPump(context)
	case *ChatMessage:
		wsa.sendMessageToClient(msg)
	}
}

func (wsa *WebSocketActor) readPump(context actor.Context) {
	defer func() {
		wsa.conn.Close()
		context.Send(wsa.chatRoom, &Unregister{Client: context.Self()})
	}()

	for {
		_, message, err := wsa.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("WebSocket read error: %v", err)
			}
			break
		}
		// Forward message to ChatRoom
		context.Send(wsa.chatRoom, &ChatMessage{Content: string(message)})
	}
}

func (wsa *WebSocketActor) sendMessageToClient(msg *ChatMessage) {
	err := wsa.conn.WriteMessage(websocket.TextMessage, []byte(msg.Content))
	if err != nil {
		log.Printf("WebSocket write error: %v", err)
	}
}
