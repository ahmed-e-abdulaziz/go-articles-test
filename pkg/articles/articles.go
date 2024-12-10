package articles

import (
	"database/sql"
	"errors"

	"github.com/ahmed-e-abdulaziz/go-articles-test/pkg/models"
	"github.com/ahmed-e-abdulaziz/go-articles-test/pkg/repository"
)

const NoArticleFoundError = "no article was found"

func GetArticleById(id int) (*models.Article, error) {
	article, err := repository.GetArticleById(id)
	if err != nil && err == sql.ErrNoRows {
		err = errors.New(NoArticleFoundError)
	}
	return article, err
}
