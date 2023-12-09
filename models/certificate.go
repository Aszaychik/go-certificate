package models

import "time"

type Certificate struct {
	ID         string
	Name       string
	CourseName string
	Date       time.Time
}