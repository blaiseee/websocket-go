package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Define a WebSocket upgrader
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Define a WebSocket handler function
func wsHandler(w http.ResponseWriter, r *http.Request) {
	// Upgrade HTTP connection to a WebSocket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading connection to WebSocket:", err)
		return
	}
	defer conn.Close()

	// Handle incoming WebSocket messages
	for {
		// Read message from client
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}

		// Print the received message
		log.Printf("Received message: %s", msg)

		// Echo the message back to the client
		if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
			log.Println("Error writing message:", err)
			break
		}
	}
}

func main() {
	// Register WebSocket handler
	http.HandleFunc("/ws", wsHandler)

	// Start the HTTP server
	log.Println("Starting server on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
