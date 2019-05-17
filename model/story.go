package model

import "time"

type Story struct {
	ID           string
	Title        string
	Content      string
	Author       string
	CreatedTime  *time.Time
	ModifiedTime *time.Time
}
