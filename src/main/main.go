package main

import "github.com/gin-gonic/gin"

//HomePage return JSON response
func HomePage(c *gin.Context) {
	c.JSON(200, gin.H{"message": "PONG"})
}

func main() {
	r := gin.Default()
	r.GET("/ping", HomePage)
	r.Run() // listen and serve on 0.0.0.0:8080
}
