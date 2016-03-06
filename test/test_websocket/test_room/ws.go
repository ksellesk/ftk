package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var hostserver = "10.14.37.137"

func main() {

	r := gin.Default()
	r.LoadHTMLFiles("index.html")

	r.GET("/chat/:roomid", chatroom)

	r.GET("/ws", func(c *gin.Context) {
		wshandler(c.Writer, c.Request)
	})

	r.Run(":8080")
}

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func chatroom(c *gin.Context) {
	roomid := c.Param("roomid")
	c.HTML(200, "index.html", gin.H{
		"roomid": roomid,
		"hostserver": hostserver,
	})

}

func wshandler(w http.ResponseWriter, r *http.Request) {
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	}

	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		conn.WriteMessage(t, msg)
	}
}
