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
	r.GET("/ws", servews)
	r.Run()
}

func index(c *gin.Context) {

	dbSet("ksel", "xuexue")
	title := "testing"
	firstget := dbGet("ksel")
	c.HTML(200, "index.html", gin.H{
		"title": title,
		"redis": firstget,
	})

}

func servews(c *gin.Context) {
	wshandler(c.Writer, c.Request)
}
