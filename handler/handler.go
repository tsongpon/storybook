package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/tsongpon/yoneebook/service"
)

// Ping Handle ping check request
func Ping(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}

// GetStory handle get story by id http request
func GetStory(c echo.Context) error {
	story, err := service.GetStory("ID")
	if err == nil {
		return c.JSON(http.StatusOK, story)
	}
	log.Print(err.Error())
	return c.String(http.StatusInternalServerError, "error")
}
