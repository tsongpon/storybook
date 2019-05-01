package service

import (
	"github.com/google/uuid"
	"github.com/tsongpon/yoneebook/model"
)

// GetStory get story by given id
func GetStory(ID string) (model.Story, error) {
	s := model.Story{
		ID:      uuid.New().String(),
		Title:   "Test title",
		Content: "Some content",
	}
	return s, nil
}
