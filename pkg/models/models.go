package models

import "time"

type Article struct {
	Id                int
	Title             string
	Content           string
	CreationTimestamp time.Time
}