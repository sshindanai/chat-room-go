package main

import "github.com/gin-gonic/gin"

// StartServer for application
var (
	router *gin.Engine
)

func main() {
	go h.run()
	router = gin.New()
	router.LoadHTMLFiles("template/index.html")

	router.GET("/room/:roomId", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	router.GET("/ws/:roomId", func(c *gin.Context) {
		roomId := c.Param("roomId")
		serveWs(c.Writer, c.Request, roomId)
	})

	router.Run("0.0.0.0:8080")
}
