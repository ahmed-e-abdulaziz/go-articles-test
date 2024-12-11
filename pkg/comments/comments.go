package comments

import (
	"errors"

	"github.com/ahmed-e-abdulaziz/go-articles-test/pkg/models"
	"github.com/ahmed-e-abdulaziz/go-articles-test/pkg/repository"
)

const NoArticleIdProvidedErrorContent = "please provide a valid ArticleId to add the comment"

var CreateComment = func(comment *models.Comment) error {
	if comment.ArticleId == 0 {
		return errors.New(NoArticleIdProvidedErrorContent)
	}
	err := repository.CreateComment(comment)
	if err != nil && err.Error() == repository.ArticleIdFKErrorContent {
		return errors.New(NoArticleIdProvidedErrorContent) // to avoid exposing the repository's error
	}
	return err
}
