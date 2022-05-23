package models

import "time"

type PersonByID struct {
	ID uint32 `json:"id" form:"id" binding:"required"`
}

type Person struct {
	Name  string    `json:"name" binding:"required"`
	Birth time.Time `json:"birth"`
}

type ParentRelationship struct {
	Parent   uint32 `json:"parent" binding:"required"`
	Children uint32 `json:"children" binding:"required"`
}

type GetPerson PersonByID
type GetAscendants PersonByID
type GetParentRelationship ParentRelationship

type PostPerson Person
type PostParentRelationship ParentRelationship

type DelPerson PersonByID
type DelParentRelationship ParentRelationship

type Request interface {
	requestBind()
}

func (req GetPerson) requestBind()             {}
func (req GetAscendants) requestBind()         {}
func (req GetParentRelationship) requestBind() {}

func (req PostPerson) requestBind()             {}
func (req PostParentRelationship) requestBind() {}

func (req DelParentRelationship) requestBind() {}

type Node struct {
	name  string
	edges []Node
}
