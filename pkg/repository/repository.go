package repository

import (
	"database/sql"
	"errors"
	"time"

	"github.com/ahmed-e-abdulaziz/go-articles-test/pkg/models"
	"github.com/jackc/pgx/v5/pgconn"
)

/*
 * A single repository struct for all DB access for ease of initialization
 * can be broken to different concrete structs in the future
 * For now it's one repository struct and multiple interfaces
 */

type Repository struct {
	db *sql.DB
}

type ArticleRepository interface {
	GetArticleById(id int) (*models.Article, error)
	GetArticles() ([]models.Article, error)
	CreateArticle(article *models.Article) error
}

type CommentRepository interface {
	CreateComment(comment *models.Comment) error
	GetCommentsByArticleId(articleId int) ([]models.Comment, error)
}

func NewRepository(db *sql.DB) *Repository {
	repo := new(Repository)
	repo.db = db
	return repo
}

const ArticleIdFKErrorContent = "foreign key constraint error occured for article id in comment creation"

func (repo *Repository) GetArticleById(id int) (*models.Article, error) {
	article := new(models.Article)
	result := repo.db.QueryRow("SELECT id, title, content, creation_timestamp FROM article WHERE ID = $1", id)
	err := result.Scan(&article.Id, &article.Title, &article.Content, &article.CreationTimestamp)
	return article, err
}

func (repo *Repository) GetArticles() ([]models.Article, error) {
	result := []models.Article{}
	rows, err := repo.db.Query("SELECT id, title, content, creation_timestamp FROM article")
	for rows.Next() {
		article := new(models.Article)
		rows.Scan(&article.Id, &article.Title, &article.Content, &article.CreationTimestamp)
		result = append(result, *article)
	}
	return result, err
}

func (repo *Repository) CreateArticle(article *models.Article) error {
	if article.CreationTimestamp.IsZero() {
		article.CreationTimestamp = time.Now()
	}
	_, err := repo.db.Exec("INSERT INTO article(title, content, creation_timestamp) VALUES ($1, $2, $3)",
		article.Title, article.Content, article.CreationTimestamp)
	return err
}

func (repo *Repository) CreateComment(comment *models.Comment) error {
	if comment.CreationTimestamp.IsZero() {
		comment.CreationTimestamp = time.Now()
	}
	_, err := repo.db.Exec("INSERT INTO comment(article_id, author, content, creation_timestamp) VALUES ($1, $2, $3, $4)",
		comment.ArticleId, comment.Author, comment.Content, comment.CreationTimestamp)
	if pgerr, ok := err.(*pgconn.PgError); ok {
		if pgerr.Code == "23503" { // FOREIGN KEY VIOLATION code in postgres
			return errors.New(ArticleIdFKErrorContent)
		}
	}
	return err
}

func (repo *Repository) GetCommentsByArticleId(articleId int) ([]models.Comment, error) {
	result := []models.Comment{}
	rows, err := repo.db.Query("SELECT id, article_id, author, content, creation_timestamp FROM comment")
	for rows.Next() {
		comment := new(models.Comment)
		rows.Scan(&comment.Id, &comment.ArticleId, &comment.Author, &comment.Content, &comment.CreationTimestamp)
		result = append(result, *comment)
	}
	return result, err
}
