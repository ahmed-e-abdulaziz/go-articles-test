package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/ahmed-e-abdulaziz/go-articles-test/pkg/articles"
	"github.com/ahmed-e-abdulaziz/go-articles-test/pkg/comments"
	"github.com/ahmed-e-abdulaziz/go-articles-test/pkg/handlers"
	"github.com/ahmed-e-abdulaziz/go-articles-test/pkg/repository"
	"github.com/ahmed-e-abdulaziz/go-articles-test/pkg/utils"
	"github.com/gin-gonic/gin"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
)

const currentApiVersionUri = "/v1"
const articlesUri = currentApiVersionUri + "/articles"
const commentsUri = articlesUri + "/:id/comments"

func main() {
	// Dependency Injection
	repository := repository.NewRepository(initDb())
	articleService := articles.NewArticleService(repository)
	commentService := comments.NewCommentService(repository)
	handler := handlers.NewRouteHandler(articleService, commentService)

	// Route Defintions
	route := gin.Default()
	route.GET(articlesUri+"/:id", handler.GetArticleById)
	route.GET(articlesUri, handler.GetArticles)
	route.POST(articlesUri, handler.CreateArticle)
	route.POST(commentsUri, handler.CreateComment)
	route.GET(commentsUri, handler.GetCommentsForArticle)
	route.Run()
}

func initDb() *sql.DB {
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
	database, err := sql.Open("pgx", uri)
	if err != nil {
		database.Close()
		panic(err)
	}

	applyMigration(database)
	return database
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
}
