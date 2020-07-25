package main

import (
	"fmt"
	socketio "github.com/googollee/go-socket.io"
	"log"
	"net/http"
)

func main() {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	server.OnConnect("/", func(so socketio.Conn) error {
		log.Println("New user")
		so.Join("chat-room")
		return nil
	})

	server.OnEvent("/", "chat message", func(s socketio.Conn, msg string) {
		fmt.Println("message:", msg)
		server.BroadcastToRoom("/", "chat-room", "server response", msg)
	})

	go server.Serve()
	defer server.Close()

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./src")))

	log.Println("Server on port 3000")

	log.Fatal(http.ListenAndServe(":3000", nil))
}
