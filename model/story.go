package model

import "time"

type Story struct {
	ID           string    `json:"id"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	Author       string    `json:"author"`
	CreatedTime  time.Time `json:"createdTime"`
	ModifiedTime time.Time `json:"modifiedTime"`
}
