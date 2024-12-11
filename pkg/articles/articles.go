package articles

import (
	"database/sql"
	"errors"

	"github.com/ahmed-e-abdulaziz/go-articles-test/pkg/models"
	"github.com/ahmed-e-abdulaziz/go-articles-test/pkg/repository"
)

const NoArticleFoundError = "no article was found"

var GetArticleById = func(id int) (*models.Article, error) {
	article, err := repository.GetArticleById(id)
	if err != nil && err == sql.ErrNoRows {
		err = errors.New(NoArticleFoundError)
	}
	return article, err
}

var GetArticles = func() ([]models.Article, error) {
	return repository.GetArticles()
}

var CreateArticle = func(article *models.Article) error {
	return repository.CreateArticle(article)
}
