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

	localAddr, err := net.ResolveUDPAddr("udp", ":8081")
	if err != nil {
		log.Fatal(err)
	}

	listener, err := net.ListenUDP("udp", localAddr)
	if err != nil {
		log.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer listener.Close()

	// TODO - get ID from lobby
	world := world.NewWorld(1, listener)
	world.Initialize()

	// TODO - handle graceful shutdown
	go world.BroadcastLoop()

	// TODO - handle graceful shutdown
	world.ControlLoop()
}
