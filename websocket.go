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
//		dbPop("test"+roomid, string(msg))
		if err != nil {
			break
		}
//		msg = []byte(dbPush("test"+roomid))
		conn.WriteMessage(t, msg)
	}
}
