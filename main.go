package main

import (
	"github.com/gin-gonic/gin" 
	"os"
)

func main() {
	port := os.Getenv("PORT")

	router := gin.Default()
	router.LoadHTMLFiles("index.html")
	router.StaticFile("/style.css", "style.css")
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	router.Run(":" + port)
}
