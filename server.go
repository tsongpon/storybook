package main

import (
	"database/sql"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"

	"github.com/tsongpon/yoneebook/repository"
	"github.com/tsongpon/yoneebook/service"
	"github.com/tsongpon/yoneebook/v1/handler"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	dbHost := getEnv("DB_HOST", "localhost")
	dbUser := getEnv("DB_USER", "root")
	dbPassword := getEnv("DB_PASSWORD", "pingu123")

	db, err := sql.Open("mysql", dbUser+":"+dbPassword+"@tcp("+dbHost+":3306)/yoneebook?multiStatements=true&parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		panic(err.Error())
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"mysql",
		driver,
	)
	if err != nil {
		panic(err.Error())
	}
	m.Steps(2)

	mysqlRepo := repository.NewMysqlRepository(db)
	// repository := repository.NewInMemoryStoryRepository()
	service := service.NewStoryService(mysqlRepo)
	handler := handler.NewStoryHandler(service)

	e := echo.New()
	e.Use(middleware.Logger())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	e.GET("/ping", handler.Ping)
	e.GET("/v1/stories/:id", handler.GetStory)
	e.GET("v1/stories", handler.GetStories)
	e.POST("/v1/stories", handler.CreateStory)

	e.Logger.Fatal(e.Start(":5000"))
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
