package handler

import (
	"net/http"

	"github.com/tsongpon/yoneebook/model"

	"github.com/labstack/echo"
	"github.com/tsongpon/yoneebook/service"
)

type StoryHandler struct {
	service *service.StoryService
}

func NewStoryHandler(service *service.StoryService) *StoryHandler {
	h := new(StoryHandler)
	h.service = service
	return h
}

// Ping Handle ping check request
func (h *StoryHandler) Ping(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}

func (h *StoryHandler) GetStories(c echo.Context) error {
	var err error
	stories, err := h.service.GetStories()
	if err != nil {
		return c.String(http.StatusInternalServerError, "error")
	}
	return c.JSON(http.StatusOK, stories)
}

// GetStory handle get story by id http request
func (h *StoryHandler) GetStory(c echo.Context) error {
	id := c.Param("id")
	var err error
	story, err := h.service.GetStory(id)
	if err != nil {
		return c.String(http.StatusInternalServerError, "error")
	}
	if story == nil {
		return c.NoContent(http.StatusNotFound)
	}
	return c.JSON(http.StatusOK, story)

}

func (h *StoryHandler) CreateStory(c echo.Context) error {
	story := new(model.Story)
	if err := c.Bind(story); err != nil {
		return err
	}
	created, _ := h.service.CreateStory(story)
	return c.JSON(http.StatusCreated, created)
}
