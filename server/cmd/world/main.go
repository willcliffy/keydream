package main

import (
	"log"
	"net"
	"os"
	"time"

	"github.com/willcliffy/keydream-server/world"
)

func main() {
	// TODO - when I set up envs, toggle this off except for local
	// This is here so that the world control loop begins after the lobby is ready
	time.Sleep(3 * time.Second)

	// TODO - get ID from lobby
	var world world.World
	world.Initialize()

	listener, err := net.Listen("tcp", "localhost:80")
	if err != nil {
		log.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer listener.Close()

	// TODO - handle graceful shutdown
	go world.BroadcastLoop()

	// TODO - handle graceful shutdown
	world.ControlLoop(listener)
}
