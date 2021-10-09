package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	http.HandleFunc("/echo", func(rw http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(rw, r, nil)

		for {
			// read message from browser
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}

			// print the message to the console
			fmt.Printf("%s sent: %v %s \n", conn.RemoteAddr(), msgType, string(msg))

			// write message back to browser
			if err = conn.WriteMessage(msgType, msg); err != nil {
				return
			}
		}
	})

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		http.ServeFile(rw, r, "websockets.html")
	})

	http.ListenAndServe(":8080", nil)
}
