package actors

import (
	"net/http"

	"github.com/asynkron/protoactor-go/actor"
)

// ChatMessage represents a chat message sent by a client
type ChatMessage struct {
	Content string
}

// Register is sent by a WebSocket actor to register with the ChatRoom
type Register struct {
	Client *actor.PID
}

// Unregister is sent by a WebSocket actor to unregister from the ChatRoom
type Unregister struct {
	Client *actor.PID
}

// Add this new message type
type HTTPConnection struct {
	W http.ResponseWriter
	R *http.Request
}

