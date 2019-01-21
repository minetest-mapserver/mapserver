package web

import (
	"mapserver/app"
	"net/http"
	"github.com/gorilla/websocket"
	"fmt"
)

type WS struct {
	ctx *app.App
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (t *WS) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	conn, _ := upgrader.Upgrade(resp, req, nil)

	for {
			// Read message from browser
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}

			// Print the message to the console
			fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

			// Write message back to browser
			if err = conn.WriteMessage(msgType, msg); err != nil {
				return
			}
		}

}
