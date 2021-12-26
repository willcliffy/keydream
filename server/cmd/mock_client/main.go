package main

import (
	"bufio"
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Fatal("* dialing:", err)
	}
	defer conn.Close()
	log.Println("* Connected to server.")

	for i := 0; i < 10; i++ {
		log.Println("* Sending to server: Hello from client")
		_, err = conn.Write([]byte("Hello from client\n"))
		if err != nil {
			log.Fatal("* writing:", err)
		}

		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatal("* reading:", err)
		}
		log.Printf("* Message from server: %v", message)

		time.Sleep(5 * time.Second)
	}

	log.Println("* Closing connection.")
	conn.Close()
}
