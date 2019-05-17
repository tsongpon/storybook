package handler

import (
	"net/http"
	"strconv"

	"github.com/tsongpon/yoneebook/query"
	"github.com/tsongpon/yoneebook/v1/mapper"
	"github.com/tsongpon/yoneebook/v1/transport"

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
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	offset, err := strconv.Atoi(c.QueryParam("offset"))
	stories, err := h.service.GetStories(query.StoryQuery{Limit: limit, Offset: offset})
	if err != nil {
		return c.String(http.StatusInternalServerError, "error")
	}
	transports := []transport.StoryTransport{}
	for _, each := range stories {
		transports = append(transports, mapper.ToTransport(each))
	}
	res := transport.ResponseTransport{}
	res.Total = len(transports)
	res.Size = len(transports)
	res.Data = transports
	return c.JSON(http.StatusOK, res)
}

// GetStory handle get story by id http request
func (h *StoryHandler) GetStory(c echo.Context) error {
	id := c.Param("id")
	var err error
	story, err := h.service.GetStory(id)
	if err != nil {
		return c.String(http.StatusInternalServerError, "error")
	}
	if story.ID == "" {
		return c.NoContent(http.StatusNotFound)
	}
	return c.JSON(http.StatusOK, mapper.ToTransport(story))

}

func (h *StoryHandler) CreateStory(c echo.Context) error {
	t := transport.StoryTransport{}
	if err := c.Bind(&t); err != nil {
		return err
	}
	created, _ := h.service.CreateStory(mapper.ToModel(t))
	return c.JSON(http.StatusCreated, mapper.ToTransport(created))
}
