package model

import "time"

type StoryViewedEvent struct {
	ID        string    `json:"ID"`
	StoryID   string    `json:"storyId"`
	Time      time.Time `json:"time"`
	UserAgent string    `json:"userAgent"`
}
