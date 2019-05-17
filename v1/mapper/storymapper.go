package mapper

import (
	"github.com/tsongpon/yoneebook/model"
	"github.com/tsongpon/yoneebook/v1/transport"
)

func ToTransport(s model.Story) transport.StoryTransport {
	t := transport.StoryTransport{}
	t.ID = s.ID
	t.Title = s.Title
	t.Author = s.Author
	t.CreatedTime = s.CreatedTime
	t.ModifiedTime = s.ModifiedTime
	t.Content = s.Content
	return t
}

func ToModel(t transport.StoryTransport) model.Story {
	s := model.Story{}
	s.ID = t.ID
	s.Title = t.Title
	s.Author = t.Author
	s.CreatedTime = t.CreatedTime
	s.ModifiedTime = t.ModifiedTime
	s.Content = t.Content
	return s
}
