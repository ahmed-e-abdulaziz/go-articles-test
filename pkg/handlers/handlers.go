package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/ahmed-e-abdulaziz/go-articles-test/pkg/articles"
	"github.com/ahmed-e-abdulaziz/go-articles-test/pkg/models"
	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	err    string
	status int
}

func IdNotFoundResponse() ErrorResponse {
	return ErrorResponse{err: "Invalid or no id was supplied for GetArticleById", status: http.StatusBadRequest}
}

func ArticleByIdError(id string) ErrorResponse {
	return ErrorResponse{err: "Encountered an error while getting articles by id: " + id, status: http.StatusBadRequest}
}

func ArticleNotFound(id string) ErrorResponse {
	return ErrorResponse{err: "No article was found for id: " + id, status: http.StatusNotFound}
}

func ArticleGetAllError() ErrorResponse {
	return ErrorResponse{err: "An error occured while getting all articles", status: http.StatusInternalServerError}
}

func ArticleBindingError() ErrorResponse {
	return ErrorResponse{err: "An error occured while reading the request body as an article", status: http.StatusBadRequest}
}

func ArticleCreationError() ErrorResponse {
	return ErrorResponse{err: "An error occured while creating an article", status: http.StatusInternalServerError}
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
			c.JSON(http.StatusNotFound, articles.NoArticleFoundError)
		} else {
			log.Printf("Encountered an error while getting articles by id: %s", idParam)
			c.JSON(http.StatusBadRequest, ArticleByIdError(idParam))
		}
		return
	}
	c.JSON(http.StatusOK, article)
}

func GetArticles(c *gin.Context) {
	articles, err := articles.GetArticles()
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusInternalServerError, ArticleGetAllError())
		return
	}
	c.JSON(http.StatusOK, articles)
}

func CreateArticle(c *gin.Context) {
	article := new(models.Article)
	err := c.BindJSON(article)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusBadRequest, ArticleBindingError())
		return
	}
	err = articles.CreateArticle(article)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusInternalServerError, ArticleCreationError())
		return
	}
	c.Status(http.StatusCreated)
}
