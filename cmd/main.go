package main

import (
	"github.com/ahmed-e-abdulaziz/go-articles-test/pkg/handlers"
	"github.com/ahmed-e-abdulaziz/go-articles-test/pkg/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	repository.InitDb()
	route := gin.Default()
	route.GET("/v1/articles/:id", handlers.GetArticleById)
	route.Run() // listen and serve on 0.0.0.0:8080
}

