package models

import "time"

type AddPerson struct {
	Name  string    `json:"name" binding:"required"`
	Birth time.Time `json:"birth"`
}

type AddParentRelationship struct {
	Parent   uint32 `json:"parent" binding:"required"`
	Children uint32 `json:"children" binding:"required"`
}

type GetAscendants struct {
	Person uint32 `form:"person" binding:"required" json:"person"`
}

type Request interface {
	requestBind()
}

func (req AddPerson) requestBind()             {}
func (req AddParentRelationship) requestBind() {}
func (req GetAscendants) requestBind()         {}

type Node struct {
	name  string
	edges []Node
}
