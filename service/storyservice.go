package service

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/tsongpon/yoneebook/model"
	"github.com/tsongpon/yoneebook/query"
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

func (servive *StoryService) GetStories(q query.StoryQuery) ([]model.Story, error) {
	var err error
	stories, err := servive.repo.GetStories(q)
	if err != nil {
		return nil, err
	}
	return stories, nil
}

// GetStory get story by given id
func (service *StoryService) GetStory(ID string) (model.Story, error) {
	s, _ := service.repo.GetStory(ID)
	return s, nil
}

func (service *StoryService) CreateStory(story model.Story) (model.Story, error) {
	story.ID = uuid.New().String()
	now := time.Now()
	if story.CreatedTime == nil {
		story.CreatedTime = &now
	}
	story.ModifiedTime = &now
	s, _ := service.repo.SaveStory(story)
	return s, nil
}

func (service *StoryService) SaveStoryViewed(storyID string, userAgent string) error {
	e := model.StoryViewedEvent{ID: uuid.New().String(), StoryID: storyID, UserAgent: userAgent, Time: time.Now()}
	err := service.repo.SaveStoryViewed(e)
	if err != nil {
		log.Printf("save event error")
	}
	return nil
}
