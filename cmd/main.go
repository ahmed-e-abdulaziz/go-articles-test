package main

import (
	"github.com/ahmed-e-abdulaziz/go-articles-test/pkg/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	route.GET("/api/", handlers.GetPostsById)
	route.Run() // listen and serve on 0.0.0.0:8080
}