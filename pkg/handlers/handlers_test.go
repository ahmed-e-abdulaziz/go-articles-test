package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
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
	defer initContext()
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
}

func TestGetArticleById(t *testing.T) {
	// Given
	defer initContext()
	articles.GetArticleById = successfulGetArticleById()
	context.AddParam("id", "1")

	// When
	GetArticleById(context)

	// Then
	expected, _ := json.Marshal(validArticle(1))
	assert.Equal(t, string(expected), string(recorder.Body.String()))
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestGetArticles(t *testing.T) {
	// Given
	defer initContext()
	articles.GetArticles = successfulGetArticles()

	// When
	GetArticles(context)

	// Then
	expected, _ := json.Marshal([]models.Article{*validArticle(1), *validArticle(2)})
	assert.Equal(t, string(expected), string(recorder.Body.String()))
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestCreateArticle(t *testing.T) {
	// Given
	defer initContext()
	calledCreateArticle := false
	articles.CreateArticle = func(article *models.Article) error {
		calledCreateArticle = true
		return nil
	}

	body, _ := json.Marshal(validArticle(1))
	context.Request = &http.Request{
		URL:  &url.URL{},
		Body: io.NopCloser(bytes.NewBuffer(body)),
	}

	// When
	CreateArticle(context)

	// Then
	assert.True(t, calledCreateArticle, "Should call articles.CreateArticle with valid request")
}

func initContext() {
	recorder = httptest.NewRecorder()
	context, _ = gin.CreateTestContext(recorder)
}

func successfulGetArticles() func() ([]models.Article, error) {
	return func() ([]models.Article, error) {
		return []models.Article{*validArticle(1), *validArticle(2)}, nil
	}
}

func successfulGetArticleById() func(id int) (*models.Article, error) {
	return func(id int) (*models.Article, error) {
		return validArticle(id), nil
	}
}

func validArticle(id int) *models.Article {
	return &models.Article{Id: id, Title: "Awesome", Content: "Awesome article is awesome", CreationTimestamp: time.UnixMilli(1733829984990)}
}
