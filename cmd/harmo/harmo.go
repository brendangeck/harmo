package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/asynkron/protoactor-go/actor"
	"harmo/internal/actors"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	system := actor.NewActorSystem()

	// Create ChatRoom actor
	chatRoomProps := actor.PropsFromProducer(actors.NewChatRoomActor)
	chatRoom := system.Root.Spawn(chatRoomProps)

	// HTTP handler for WebSocket connections
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Printf("WebSocket upgrade failed: %v", err)
			return
		}

		props := actor.PropsFromProducer(actors.NewWebSocketActor(conn, chatRoom))
		system.Root.Spawn(props)
	})

	log.Println("Starting WebSocket server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
