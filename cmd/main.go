package main

import (
	"github.com/ahmed-e-abdulaziz/go-articles-test/pkg/handlers"
	"github.com/ahmed-e-abdulaziz/go-articles-test/pkg/repository"
	"github.com/gin-gonic/gin"
)

const currentApiVersionUri = "/v1"
const articlesUri = currentApiVersionUri + "/articles"
const commentsUri = articlesUri + "/:id/comments"

func main() {
	repository.InitDb()
	route := gin.Default()
	route.GET(articlesUri+"/:id", handlers.GetArticleById)
	route.GET(articlesUri, handlers.GetArticles)
	route.POST(articlesUri, handlers.CreateArticle)
	route.POST(commentsUri, handlers.CreateComment)
	route.GET(commentsUri, handlers.GetCommentsForArticle)
	route.Run()
}
