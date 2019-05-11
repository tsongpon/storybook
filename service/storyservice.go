package service

import (
	"github.com/google/uuid"
	"github.com/tsongpon/yoneebook/model"
	"github.com/tsongpon/yoneebook/repository"
)

type StoryService struct {
	repo repository.StoryRepository
}

func NewStoryService(repo repository.StoryRepository) *StoryService {
	service := new(StoryService)
	service.repo = repo
	return service
}

func (servive *StoryService) GetStories() ([]*model.Story, error) {
	var err error
	stories, err := servive.repo.GetStories()
	if err != nil {
		return nil, err
	}
	return stories, nil
}

// GetStory get story by given id
func (service *StoryService) GetStory(ID string) (*model.Story, error) {
	s, _ := service.repo.GetStory(ID)
	return s, nil
}

func (service *StoryService) CreateStory(story *model.Story) (*model.Story, error) {
	story.ID = uuid.New().String()
	s, _ := service.repo.SaveStory(story)
	return s, nil
}
