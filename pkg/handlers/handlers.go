package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/ahmed-e-abdulaziz/go-articles-test/pkg/articles"
	"github.com/ahmed-e-abdulaziz/go-articles-test/pkg/comments"
	"github.com/ahmed-e-abdulaziz/go-articles-test/pkg/handlers/errres"
	"github.com/ahmed-e-abdulaziz/go-articles-test/pkg/models"
	"github.com/gin-gonic/gin"
)

type RouteHandler struct {
	articleService articles.ArticleService
	commentService comments.CommentService
}

func NewRouteHandler(
	articleService articles.ArticleService,
	commentService comments.CommentService) *RouteHandler {
	return &RouteHandler{articleService: articleService, commentService: commentService}
}

func (h *RouteHandler) GetArticleById(c *gin.Context) {
	idParam, ok := c.Params.Get("id")
	id, err := strconv.Atoi(idParam)
	if idParam == "" || !ok || err != nil {
		log.Printf("No id was found for GetArticleById")
		c.JSON(http.StatusBadRequest, errres.ArticleIdNotFoundResponse())
		return
	}
	article, err := h.articleService.GetArticleById(id)
	if err != nil {
		if err.Error() == articles.NoArticleFoundError {
			log.Printf("No article was found for id: %s", idParam)
			c.JSON(http.StatusNotFound, articles.NoArticleFoundError)
		} else {
			log.Printf("Encountered an error while getting articles by id: %s", idParam)
			c.JSON(http.StatusBadRequest, errres.ArticleByIdError(idParam))
		}
		return
	}
	c.JSON(http.StatusOK, article)
}

func (h *RouteHandler) GetArticles(c *gin.Context) {
	articles, err := h.articleService.GetArticles()
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusInternalServerError, errres.ArticleGetAllError())
		return
	}
	c.JSON(http.StatusOK, articles)
}

func (h *RouteHandler) CreateArticle(c *gin.Context) {
	article := new(models.Article)
	err := c.BindJSON(article)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusBadRequest, errres.ArticleBindingError())
		return
	}
	err = h.articleService.CreateArticle(article)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusInternalServerError, errres.ArticleCreationError())
		return
	}
	c.Status(http.StatusCreated)
}

func (h *RouteHandler) CreateComment(c *gin.Context) {
	idParam, ok := c.Params.Get("id")
	articleId, err := strconv.Atoi(idParam)
	if idParam == "" || !ok || err != nil {
		log.Printf("No id was provided for CreateComment")
		c.JSON(http.StatusBadRequest, errres.ArticleIdNotFoundResponse())
		return
	}
	comment := new(models.Comment)
	err = c.BindJSON(comment)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusBadRequest, errres.CommentBindingError())
		return
	}
	comment.ArticleId = articleId
	err = h.commentService.CreateComment(comment)
	if err != nil {
		if err.Error() == comments.NoArticleIdProvidedErrorContent {
			log.Print(err.Error())
			c.JSON(http.StatusBadRequest, errres.CommentInvalidArticleIdProvidedError())
		}
		log.Print(err.Error())
		c.JSON(http.StatusInternalServerError, errres.CommentCreationError())
		return
	}
	c.Status(http.StatusCreated)
}

func (h *RouteHandler) GetCommentsForArticle(c *gin.Context) {
	idParam, ok := c.Params.Get("id")
	articleId, err := strconv.Atoi(idParam)
	if idParam == "" || !ok || err != nil {
		log.Printf("No id was provided for CreateComment")
		c.JSON(http.StatusBadRequest, errres.ArticleIdNotFoundResponse())
		return
	}
	comments, err := h.commentService.GetCommentsByArticleId(articleId)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusInternalServerError, errres.ArticleGetAllError())
		return
	}
	c.JSON(http.StatusOK, comments)
}
