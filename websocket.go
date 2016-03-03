package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func wshandler(w http.ResponseWriter, r *http.Request, roomid string) {
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	}

	for {
		t, msg, err := conn.ReadMessage()
//		dbSet("test"+roomid, string(msg))
		if err != nil {
			break
		}
//		msg = []byte(dbGet("test"+roomid))
		conn.WriteMessage(t, msg)
	}
}
