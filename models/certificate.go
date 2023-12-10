package models

import "time"

type Certificate struct {
	ID         string
	Name       string
	CourseName string
	CreatedAt  time.Time
	IssuedAt   string
}