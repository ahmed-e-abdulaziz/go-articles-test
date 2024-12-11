package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/ahmed-e-abdulaziz/go-articles-test/pkg/models"
	"github.com/ahmed-e-abdulaziz/go-articles-test/pkg/utils"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgconn"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var db *sql.DB

const ArticleIdFKErrorContent = "foreign key constraint error occured for article id in comment creation"

func InitDb() {
	driverName := os.Getenv("DATABASE_DRIVER") // e.g. postgres
	username := os.Getenv("DATABASE_USERNAME") // e.g. postgres
	password := os.Getenv("DATABASE_PASSWORD") // e.g. password
	host := os.Getenv("DATABASE_HOST")         // e.g. localhost
	port := os.Getenv("DATABASE_PORT")         // e.g. 5432
	if utils.ContainsEmpty(driverName, username, password, host, port) {
		panicMessage := fmt.Sprintf(`Please provide the following environment variables before starting:
			- DATABASE_DRIVER (e.g. postgres) the provided value was [%s]
			- DATABASE_USERNAME (e.g. postgres) the provided value was [%s]
			- DATABASE_PASSWORD (e.g. password) the provided value was [%s]
			- DATABASE_HOST (e.g. localhost) the provided value was [%s]
			- DATABASE_PORT (e.g. 5432) the provided value was [%s]`,
			driverName, username, password, host, port)
		panic(panicMessage)
	}
	uri := driverName + "://" + username + ":" + password + "@" + host + ":" + port + "/articles"
	database, err := sql.Open("pgx", uri) //TODO: Add support for MySQL lib, make pgx an env var
	if err != nil {
		database.Close()
		panic(err)
	}

	applyMigration(database)
}

func applyMigration(database *sql.DB) {
	migrationDriver, err := postgres.WithInstance(database, &postgres.Config{})
	if err != nil {
		panic(err)
	}
	m, err := migrate.NewWithDatabaseInstance("file://db", "articles", migrationDriver)
	if err != nil {
		panic(err)
	}
	m.Up()
	db = database
}

func GetArticleById(id int) (*models.Article, error) {
	article := new(models.Article)
	result := db.QueryRow("SELECT id, title, content, creation_timestamp FROM article WHERE ID = $1", id)
	err := result.Scan(&article.Id, &article.Title, &article.Content, &article.CreationTimestamp)
	return article, err
}

func GetArticles() ([]models.Article, error) {
	result := []models.Article{}
	rows, err := db.Query("SELECT id, title, content, creation_timestamp FROM article")
	for rows.Next() {
		article := new(models.Article)
		rows.Scan(&article.Id, &article.Title, &article.Content, &article.CreationTimestamp)
		result = append(result, *article)
	}
	return result, err
}

func CreateArticle(article *models.Article) error {
	if article.CreationTimestamp.IsZero() {
		article.CreationTimestamp = time.Now()
	}
	_, err := db.Exec("INSERT INTO article(title, content, creation_timestamp) VALUES ($1, $2, $3)",
		article.Title, article.Content, article.CreationTimestamp)
	return err
}

func CreateComment(comment *models.Comment) error {
	if comment.CreationTimestamp.IsZero() {
		comment.CreationTimestamp = time.Now()
	}
	_, err := db.Exec("INSERT INTO comment(article_id, author, content, creation_timestamp) VALUES ($1, $2, $3, $4)",
		comment.ArticleId, comment.Author, comment.Content, comment.CreationTimestamp)
	if pgerr, ok := err.(*pgconn.PgError); ok {
		if pgerr.Code == "23503" { // FOREIGN KEY VIOLATION code in postgres
			return errors.New(ArticleIdFKErrorContent)
		}
	}
	return err
}

func GetCommentsByArticleId(articleId int) ([]models.Comment, error) {
	result := []models.Comment{}
	rows, err := db.Query("SELECT id, article_id, author, content, creation_timestamp FROM comment")
	for rows.Next() {
		comment := new(models.Comment)
		rows.Scan(&comment.Id, &comment.ArticleId, &comment.Author, &comment.Content, &comment.CreationTimestamp)
		result = append(result, *comment)
	}
	return result, err
}
