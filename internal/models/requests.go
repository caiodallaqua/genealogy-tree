package models

import "time"

// ----------------------      BASE      ----------------------

type PersonByID struct {
	ID uint32 `json:"id" xml:"id" yaml:"id" form:"id" binding:"required" example:"42"`
}

type Person struct {
	Name  string    `json:"name" binding:"required" example:"Aristeu"`
	Birth time.Time `json:"birth" example:"2020-12-09T16:09:53+00:00"`
}

type ParentRelationship struct {
	ParentID uint32 `json:"parent_id" yaml:"parent_id" binding:"required" example:"1"`
	ChildID  uint32 `json:"child_id" yaml:"child_id" binding:"required" example:"2"`
}

// ----------------------      DERIVED      ----------------------

type GetPerson PersonByID
type GetAscendants PersonByID
type GetAscendantsAndDescendants PersonByID
type GetAscendantsAndChildren PersonByID
type GetParentRelationship ParentRelationship

type PostPerson Person
type PostParentRelationship ParentRelationship

type DelPerson PersonByID
type DelParentRelationship ParentRelationship

// To avoid repeating code that should work for every request model
type Request interface {
	GetPerson |
		GetAscendants |
		GetAscendantsAndDescendants |
		GetAscendantsAndChildren |
		GetParentRelationship |

		PostPerson |
		PostParentRelationship |

		DelPerson |
		DelParentRelationship
}
