package main

import (
	"log"
	"time"

	"github.com/asynkron/protoactor-go/actor"
	"harmo/internal/actors" 
)

func main() {
	// Create an actor system
	system := actor.NewActorSystem()

	// Create props for our actor
	props := actor.PropsFromProducer(actors.NewHelloActor)

	// Spawn the actor
	pid := system.Root.Spawn(props)

	log.Println("Sending Hello message to actor")
	// Send a message to the actor
	system.Root.Send(pid, &actors.Hello{Who: "World"})

	// Wait a bit for the message to be processed
	time.Sleep(time.Second)

	log.Println("Stopping the actor")
	// Stop the actor
	system.Root.Stop(pid)

	log.Println("Program finished")
}
