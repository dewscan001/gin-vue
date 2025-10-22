package main

import (
	"github.com/gin-gonic/gin" 
)

func testPrint(c *gin.Context){
	c.JSON(200, gin.H{"test":"test"})
}

func test(r *gin.RouterGroup) {
	r.GET("/", testPrint)
}