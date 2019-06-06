package repository

import (
	"github.com/tsongpon/yoneebook/model"
	"github.com/tsongpon/yoneebook/query"
)

type StoryRepository interface {
	GetStory(string) (model.Story, error)
	SaveStory(model.Story) (model.Story, error)
	GetStories(query.StoryQuery) ([]model.Story, error)
	SaveStoryViewed(model.StoryViewedEvent) error
}
