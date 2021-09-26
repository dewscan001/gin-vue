package main

import (
	"github.com/gin-gonic/gin" 
	"os"
)

type score_type struct{
	score int
	point_board point_type
}

type point_type struct{
	x int
	y int
}


func main() {
	port := os.Getenv("PORT")

	router := gin.Default()
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context){
		c.HTML(200, "index.tmpl", nil)
	})

	tic_tac_toe_router(router.Group("/tic-tac-toe"))

	if port != ""{
		gin.SetMode(gin.ReleaseMode)
		router.Run(":" + port)
	} else{
		router.Run(":8080")
	}
}