package repository

import (
	"github.com/tsongpon/yoneebook/model"
)

type StoryRepository interface {
	GetStory(ID string) (*model.Story, error)
	SaveStory(story *model.Story) (*model.Story, error)
	GetStories() ([]*model.Story, error)
}
