package actors

import (
	"log"

	"github.com/asynkron/protoactor-go/actor"
)

// Hello is a message type
type Hello struct{ Who string }

// HelloActor is an actor that responds to Hello messages
type HelloActor struct{}

// Receive handles incoming messages for HelloActor
func (state *HelloActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *Hello:
		log.Printf("HelloActor received Hello message with Who = %v\n", msg.Who)
	}
}

// NewHelloActor creates a new HelloActor
func NewHelloActor() actor.Actor {
	return &HelloActor{}
}
