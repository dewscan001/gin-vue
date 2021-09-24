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
		go func() {
			c.HTML(200, "index.html", nil)
		}()
	})
	if(port == ""){
		router.Run(":" + port)
	} else{
		router.Run(":8080")
	}
}
