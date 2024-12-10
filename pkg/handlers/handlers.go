package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/ahmed-e-abdulaziz/go-articles-test/pkg/articles"
	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	err    string
	status int
}

func IdNotFoundResponse() ErrorResponse {
	return ErrorResponse{err: "Invalid or no id was supplied for GetArticleById", status: http.StatusBadRequest}
}

func ArticleError(id string) ErrorResponse {
	return ErrorResponse{err: "Encountered an error while getting articles by id: " + id, status: http.StatusBadRequest}
}

func ArticleNotFound(id string) ErrorResponse {
	return ErrorResponse{err: "No article was found for id: " + id, status: http.StatusBadRequest}
}

func GetArticleById(c *gin.Context) {
	idParam, ok := c.Params.Get("id")
	id, err := strconv.Atoi(idParam)
	if idParam == "" || !ok || err != nil {
		log.Printf("No id was found for GetArticleById")
		c.JSON(http.StatusBadRequest, IdNotFoundResponse())
		return
	}
	article, err := articles.GetArticleById(id)
	if err != nil {
		if err.Error() == articles.NoArticleFoundError {
			log.Printf("No article was found for id: %s", idParam)
		} else {
			log.Printf("Encountered an error while getting articles by id: %s", idParam)
			c.JSON(http.StatusBadRequest, ArticleError(idParam))
		}
		return
	}
	c.JSON(http.StatusOK, article)
}
