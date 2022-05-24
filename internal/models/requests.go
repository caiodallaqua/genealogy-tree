package models

import "time"

type PersonByID struct {
	ID uint32 `json:"id" xml:"id" yaml:"id" form:"id" binding:"required"`
}

type Person struct {
	Name  string    `json:"name" binding:"required"`
	Birth time.Time `json:"birth"`
}

type ParentRelationship struct {
	ParentID uint32 `json:"parent_id" yaml:"parent_id" binding:"required"`
	ChildID  uint32 `json:"child_id" yaml:"child_id" binding:"required"`
}

type GetPerson PersonByID
type GetAscendants PersonByID
type GetAscendantsAndDescendants PersonByID
type GetAscendantsAndChildren PersonByID
type GetParentRelationship ParentRelationship

type PostPerson Person
type PostParentRelationship ParentRelationship

type DelPerson PersonByID
type DelParentRelationship ParentRelationship

type Request interface {
	GetPerson | GetAscendants | GetAscendantsAndDescendants | GetAscendantsAndChildren | GetParentRelationship |
		PostPerson | PostParentRelationship |
		DelPerson | DelParentRelationship
}
