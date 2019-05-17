package transport

import "time"

type ResponseTransport struct {
	Total int              `json:"total"`
	Size  int              `json:"size"`
	Data  []StoryTransport `json:"data"`
}

type StoryTransport struct {
	ID           string     `json:"id"`
	Title        string     `json:"title"`
	Content      string     `json:"content"`
	Author       string     `json:"author"`
	CreatedTime  *time.Time `json:"createdTime"`
	ModifiedTime *time.Time `json:"modifiedTime"`
}
