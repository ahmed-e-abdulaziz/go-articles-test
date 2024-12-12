package comments

import (
	"errors"

	"github.com/ahmed-e-abdulaziz/go-articles-test/pkg/models"
	"github.com/ahmed-e-abdulaziz/go-articles-test/pkg/repository"
)

type commentService struct {
	repo repository.CommentRepository
}

type CommentService interface {
	CreateComment(comment *models.Comment) error
	GetCommentsByArticleId(articleId int) ([]models.Comment, error)
}

func NewCommentService(repo *repository.Repository) CommentService {
	return &commentService{repo: repo}
}

const NoArticleIdProvidedErrorContent = "please provide a valid ArticleId to add the comment"

func (service *commentService) CreateComment(comment *models.Comment) error {
	if comment.ArticleId == 0 {
		return errors.New(NoArticleIdProvidedErrorContent)
	}
	err := service.repo.CreateComment(comment)
	if err != nil && err.Error() == repository.ArticleIdFKErrorContent {
		return errors.New(NoArticleIdProvidedErrorContent) // to avoid exposing the repository's error
	}
	return err
}

func (service *commentService) GetCommentsByArticleId(articleId int) ([]models.Comment, error) {
	return service.repo.GetCommentsByArticleId(articleId)
}
