package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
)

var connections = make([]net.Conn, 0)
var connMutex sync.Mutex

func main() {
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println("Error setting up listener:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Server is listening on port 8080...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		connMutex.Lock()
		connections = append(connections, conn)
		connMutex.Unlock()
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer func() {
		conn.Close()
		connMutex.Lock()
		for i, c := range connections {
			if c == conn {
				connections = append(connections[:i], connections[i+1:]...)
				break
			}
		}
		connMutex.Unlock()
	}()

	fmt.Println("Client connected:", conn.RemoteAddr().String())

	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading message:", err)
			return
		}

		fmt.Print("Received message:", string(message))
		broadcastMessage(conn, message)
	}
}

func broadcastMessage(sender net.Conn, message string) {
	connMutex.Lock()
	defer connMutex.Unlock()
	for _, conn := range connections {
		if conn != sender {
			_, err := conn.Write([]byte(strings.ToUpper(message) + " From Server!\n"))
			if err != nil {
				fmt.Println("Error sending message to", conn.RemoteAddr().String(), ":", err)
			}
		}
	}
}
