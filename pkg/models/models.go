package models

import "time"

type Article struct {
	Id                int       `json:"id"`
	Title             string    `json:"title"`
	Content           string    `json:"content"`
	CreationTimestamp time.Time `json:"creation_timestamp"`
}
