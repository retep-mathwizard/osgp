package main

import (
	"net/http"

	"golang.org/x/net/websocket"
)

func main() {
	http.Handle("/", websocket.Handler(func(ws *websocket.Conn) {
		ws.Write([]byte("Testing"))
	}))
	http.ListenAndServe(":3000", nil)
}
