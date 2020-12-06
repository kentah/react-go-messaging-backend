package main

import (
	"fmt"
	"net/http"

	"github.com/kentah/realtime-chat-go-react/pkg/websocket"
)

// define the websocket endpoint
func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("WebSocket Endpoint Hit")

	// upgrade this connection to a websocket connection
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprint(w, "%+v\n", err)
	}

	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}

func setupRoutes() {
	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})
}

func main() {
	setupRoutes()
	fmt.Println("web server started on port 8080")
	http.ListenAndServe(":8080", nil)
}
