package articles

import (
	"database/sql"
	"errors"

	"github.com/ahmed-e-abdulaziz/go-articles-test/pkg/models"
	"github.com/ahmed-e-abdulaziz/go-articles-test/pkg/repository"
)

type articleService struct {
	repo repository.ArticleRepository
}

type ArticleService interface {
	GetArticleById(id int) (*models.Article, error)
	GetArticles() ([]models.Article, error)
	CreateArticle(article *models.Article) error
}

func NewArticleService(repo repository.ArticleRepository) ArticleService {
	return &articleService{repo: repo}
}

const NoArticleFoundError = "no article was found"

func (service *articleService) GetArticleById(id int) (*models.Article, error) {
	article, err := service.repo.GetArticleById(id)
	if err != nil && err == sql.ErrNoRows {
		err = errors.New(NoArticleFoundError)
	}
	return article, err
}

func (service *articleService) GetArticles() ([]models.Article, error) {
	return service.repo.GetArticles()
}

func (service *articleService) CreateArticle(article *models.Article) error {
	return service.repo.CreateArticle(article)
}
