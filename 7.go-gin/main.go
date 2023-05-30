package main

import (
	"go-gin/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.GET("/posts/:id", handler.GetPostHandler)
	router.Run(":8080")
}
