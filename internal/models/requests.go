package models

import "time"

type NewPerson struct {
	Id       string    `json:"id"`
	Name     string    `json:"name" binding:"required"`
	Birth    time.Time `json:"birth" binding:"required"`
	ChildOf  []string  `json:"child_of"`
	ParentOf []string  `json:"parent_of"`
}
