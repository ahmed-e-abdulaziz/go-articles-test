package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ahmed-e-abdulaziz/go-articles-test/pkg/articles"
	"github.com/ahmed-e-abdulaziz/go-articles-test/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var recorder *httptest.ResponseRecorder
var context *gin.Context

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	initContext()
	m.Run()
}

func TestGetArticleByIdShouldReturn400(t *testing.T) {
	t.Run("No ID provided", func(t *testing.T) {
		GetArticleById(context) // No ID param provided
		assert.Equal(t, http.StatusBadRequest, recorder.Code, "When calling GetArticleById without an ID it must return 400 error")
	})
	t.Run("Empty ID provided", func(t *testing.T) {
		context.AddParam("id", "")
		GetArticleById(context) // No ID param provided
		assert.Equal(t, http.StatusBadRequest, recorder.Code, "When calling GetArticleById with an empty ID it must return 400 error")
	})
	t.Run("Non-numeric ID provided", func(t *testing.T) {
		context.AddParam("id", "ABC")
		GetArticleById(context) // No ID param provided
		assert.Equal(t, http.StatusBadRequest, recorder.Code, "When calling GetArticleById with an non numeric ID it must return 400 error")
	})
	// Cleanup
	initContext()
}

func TestGetArticleById(t *testing.T) {
	// Given
	defer initContext()
	articles.GetArticleById = successfulGetArticleById()
	context.AddParam("id", "1")

	// When
	GetArticleById(context)

	// Then
	expected, _ := json.Marshal(validArticle())
	assert.Equal(t, string(expected), string(recorder.Body.String()))
}

func initContext() {
	recorder = httptest.NewRecorder()
	context, _ = gin.CreateTestContext(recorder)
}

func successfulGetArticleById() func(id int) (*models.Article, error) {
	return func(id int) (*models.Article, error) {
		return validArticle(), nil
	}
}

func validArticle() *models.Article {
	return &models.Article{Id: 1, Title: "Awesome", Content: "Awesome article is awesome", CreationTimestamp: time.UnixMilli(1733829984990)}
}
