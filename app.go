package main

import (
	"crud-go/pet"
	"github.com/gin-gonic/gin"
)

import (
	database "crud-go/configs"
)

func main() {
	r := gin.Default()
	database.ConnectDB()

	v1 := r.Group("/v1")
	{
		pet.Controller(v1)
	}
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
