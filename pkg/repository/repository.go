package repository

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/ahmed-e-abdulaziz/go-articles-test/pkg/models"
	"github.com/ahmed-e-abdulaziz/go-articles-test/pkg/utils"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var db *sql.DB

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
