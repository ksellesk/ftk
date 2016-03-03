package main

import "github.com/gin-gonic/gin"

func main() {
	dbConn()
	ws()
	dbClose()
}

func ws() {
	r := gin.Default()
	r.LoadHTMLGlob("webroot/templates/*.html")
	r.Static("/static", "webroot/static")
	r.GET("/", index)
	r.GET("/chat/:roomid", chatroom)
	r.GET("/ws/:roomid", servews)
	r.Run()
}

func index(c *gin.Context) {

	dbSet("ksel", "xuexue")
	title := "testing"
	firstget := dbGet("ksel")
	c.HTML(200, "index.html", gin.H{
		"title": title,
		"redis": firstget,
		"roomid": "nearest",
	})

}

func chatroom(c *gin.Context) {

	roomid := c.Param("roomid")
	c.HTML(200, "index.html", gin.H{
		"title": "chatroom:"+roomid,
		"redis": "test",
		"roomid": roomid,
	})

}


func servews(c *gin.Context) {
	roomid := c.Param("roomid")
	wshandler(c.Writer, c.Request, roomid)
}
