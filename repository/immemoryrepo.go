package repository

import "github.com/tsongpon/yoneebook/model"

type InMemoryStoryRepository struct {
	storage map[string]model.Story
}

func NewInMemoryStoryRepository() *InMemoryStoryRepository {
	repo := new(InMemoryStoryRepository)
	repo.storage = make(map[string]model.Story)
	return repo
}

func (repo *InMemoryStoryRepository) GetStories() ([]*model.Story, error) {
	result := []*model.Story{}
	for _, val := range repo.storage {
		result = append(result, &val)
	}
	return result, nil
}

// GetStory return story by given ID
func (repo *InMemoryStoryRepository) GetStory(ID string) (*model.Story, error) {
	if s, ok := repo.storage[ID]; ok {
		return &s, nil
	}
	return nil, nil
}

func (repo *InMemoryStoryRepository) SaveStory(story *model.Story) (*model.Story, error) {
	repo.storage[story.ID] = *story
	return story, nil
}
