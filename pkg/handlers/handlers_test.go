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

	"github.com/ahmed-e-abdulaziz/go-articles-test/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var recorder *httptest.ResponseRecorder
var context *gin.Context

type mockArticleService struct {
	CreateArticleCalled bool
}
type mockCommentService struct {
	CalledCreateComment bool
}

var routeHandler = &RouteHandler{articleService: &mockArticleService{}, commentService: &mockCommentService{}}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	initContext()
	m.Run()
}

func TestGetArticleByIdShouldReturn400(t *testing.T) {
	defer initContext()
	t.Run("No ID provided", func(t *testing.T) {
		routeHandler.GetArticleById(context) // No ID param provided
		assert.Equal(t, http.StatusBadRequest, recorder.Code, "When calling GetArticleById without an ID it must return 400 error")
	})
	t.Run("Empty ID provided", func(t *testing.T) {
		context.AddParam("id", "")
		routeHandler.GetArticleById(context) // No ID param provided
		assert.Equal(t, http.StatusBadRequest, recorder.Code, "When calling GetArticleById with an empty ID it must return 400 error")
	})
	t.Run("Non-numeric ID provided", func(t *testing.T) {
		context.AddParam("id", "ABC")
		routeHandler.GetArticleById(context) // No ID param provided
		assert.Equal(t, http.StatusBadRequest, recorder.Code, "When calling GetArticleById with an non numeric ID it must return 400 error")
	})
}

func TestGetArticleById(t *testing.T) {
	// Given
	defer initContext()
	context.AddParam("id", "1")

	// When
	routeHandler.GetArticleById(context)

	// Then
	expected, _ := json.Marshal(validArticle(1))
	assert.Equal(t, string(expected), string(recorder.Body.String()))
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestGetArticles(t *testing.T) {
	// Given
	defer initContext()

	// When
	routeHandler.GetArticles(context)

	// Then
	expected, _ := json.Marshal([]models.Article{*validArticle(1), *validArticle(2)})
	assert.Equal(t, string(expected), string(recorder.Body.String()))
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestCreateArticle(t *testing.T) {
	// Given
	defer initContext()
	defer routeHandler.articleService.(*mockArticleService).Reset()

	body, _ := json.Marshal(validArticle(1))
	context.Request = &http.Request{
		URL:  &url.URL{},
		Body: io.NopCloser(bytes.NewBuffer(body)),
	}

	// When
	routeHandler.CreateArticle(context)

	// Then
	assert.True(t, routeHandler.articleService.(*mockArticleService).CreateArticleCalled, "Should call articleService.CreateArticle with a valid request")
}

func TestCreateCommentShouldReturn400ForWrongId(t *testing.T) {
	defer initContext()
	t.Run("No ID provided", func(t *testing.T) {
		routeHandler.CreateComment(context) // No ID param provided
		assert.Equal(t, http.StatusBadRequest, recorder.Code, "When calling GetArticleById without an ID it must return 400 error")
	})
	t.Run("Empty ID provided", func(t *testing.T) {
		context.AddParam("id", "")
		routeHandler.CreateComment(context) // No ID param provided
		assert.Equal(t, http.StatusBadRequest, recorder.Code, "When calling GetArticleById with an empty ID it must return 400 error")
	})
	t.Run("Non-numeric ID provided", func(t *testing.T) {
		context.AddParam("id", "ABC")
		routeHandler.CreateComment(context) // No ID param provided
		assert.Equal(t, http.StatusBadRequest, recorder.Code, "When calling GetArticleById with an non numeric ID it must return 400 error")
	})
}

func TestCreateComment(t *testing.T) {
	// Given
	defer initContext()
	defer routeHandler.commentService.(*mockCommentService).Reset()

	body, _ := json.Marshal(validComment(1))
	context.Request = &http.Request{
		URL:  &url.URL{},
		Body: io.NopCloser(bytes.NewBuffer(body)),
	}
	context.AddParam("id", "1")

	// When
	routeHandler.CreateComment(context)

	// Then
	assert.True(t, routeHandler.commentService.(*mockCommentService).CalledCreateComment, "Should call comments.CreateComment with a valid request")
}

func TestGetCommentsForArticleShouldReturn400ForWrongId(t *testing.T) {
	defer initContext()
	t.Run("No ID provided", func(t *testing.T) {
		routeHandler.GetCommentsForArticle(context) // No ID param provided
		assert.Equal(t, http.StatusBadRequest, recorder.Code, "When calling GetArticleById without an ID it must return 400 error")
	})
	t.Run("Empty ID provided", func(t *testing.T) {
		context.AddParam("id", "")
		routeHandler.GetCommentsForArticle(context) // No ID param provided
		assert.Equal(t, http.StatusBadRequest, recorder.Code, "When calling GetArticleById with an empty ID it must return 400 error")
	})
	t.Run("Non-numeric ID provided", func(t *testing.T) {
		context.AddParam("id", "ABC")
		routeHandler.GetCommentsForArticle(context) // No ID param provided
		assert.Equal(t, http.StatusBadRequest, recorder.Code, "When calling GetArticleById with an non numeric ID it must return 400 error")
	})
}

func TestGetCommentsForArticle(t *testing.T) {
	// Given
	defer initContext()
	context.AddParam("id", "1")

	// When
	routeHandler.GetCommentsForArticle(context)

	// Then
	expected, _ := json.Marshal([]models.Comment{*validComment(1), *validComment(2)})
	assert.Equal(t, string(expected), string(recorder.Body.String()))
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func initContext() {
	recorder = httptest.NewRecorder()
	context, _ = gin.CreateTestContext(recorder)
}

func (m *mockArticleService) GetArticleById(id int) (*models.Article, error) {
	return validArticle(id), nil
}

func (m *mockArticleService) GetArticles() ([]models.Article, error) {
	return []models.Article{*validArticle(1), *validArticle(2)}, nil
}

func (m *mockArticleService) CreateArticle(article *models.Article) error {
	m.CreateArticleCalled = true
	return nil
}

func (m *mockArticleService) Reset() {
	m.CreateArticleCalled = false
}

func (m *mockCommentService) GetCommentsByArticleId(articleId int) ([]models.Comment, error) {
	return []models.Comment{*validComment(1), *validComment(2)}, nil
}

func (m *mockCommentService) CreateComment(article *models.Comment) error {
	m.CalledCreateComment = true
	return nil
}

func (m *mockCommentService) Reset() {
	m.CalledCreateComment = false
}

func validArticle(id int) *models.Article {
	return &models.Article{Id: id, Title: "Awesome", Content: "Awesome article is awesome", CreationTimestamp: time.UnixMilli(1733829984990)}
}

func validComment(id int) *models.Comment {
	return &models.Comment{Id: id, ArticleId: 1, Author: "Ahmed Ehab", Content: "I like this awesome project and article", CreationTimestamp: time.UnixMilli(1733829984990)}
}
