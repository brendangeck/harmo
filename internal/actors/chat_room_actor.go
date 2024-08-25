package actors

import (
	"log"

	"github.com/asynkron/protoactor-go/actor"
)

type ChatRoomActor struct {
	clients map[*actor.PID]bool
}

func NewChatRoomActor() actor.Actor {
	return &ChatRoomActor{
		clients: make(map[*actor.PID]bool),
	}
}

func (cra *ChatRoomActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *actor.Started:
		log.Println("ChatRoom actor started")
	case *Register:
		cra.handleRegister(msg)
	case *Unregister:
		cra.handleUnregister(msg)
	case *ChatMessage:
		cra.handleChatMessage(msg, context)
	}
}

func (cra *ChatRoomActor) handleRegister(msg *Register) {
	cra.clients[msg.Client] = true
	log.Printf("Client registered: %v", msg.Client)
}

func (cra *ChatRoomActor) handleUnregister(msg *Unregister) {
	if _, ok := cra.clients[msg.Client]; ok {
		delete(cra.clients, msg.Client)
		log.Printf("Client unregistered: %v", msg.Client)
	}
}

func (cra *ChatRoomActor) handleChatMessage(msg *ChatMessage, context actor.Context) {
	log.Printf("Broadcasting message: %s", msg.Content)
	for client := range cra.clients {
		context.Send(client, msg)
	}
}
