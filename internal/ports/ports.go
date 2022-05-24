package ports

import (
	"genealogy-tree/internal/models"

	"github.com/gin-gonic/gin"
)

// Driver

type RestPort interface {
	CreateRouter() *gin.Engine
	RunRestServer(*gin.Engine) error
}

// API
// Connects Rest to DB
type APIPort interface {
	// GET
	CallGetStatus() error
	CallGetPerson(models.GetPerson) (any, error)
	CallGetAscendants(models.GetAscendants) (any, error)
	CallGetAscendantsAndDescendants(models.GetAscendantsAndDescendants) (any, error)
	CallGetAscendantsAndChildren(models.GetAscendantsAndChildren) (any, error)

	CallPostPerson(models.PostPerson) (any, error)
	CallPostParentRelationship(models.PostParentRelationship) (any, error)

	CallDelPerson(models.DelPerson) (any, error)
	CallDelParentRelationship(models.DelParentRelationship) (any, error)
}

// Driven

type DBPort interface {
	// GET
	GetStatus() error
	GetPerson(models.GetPerson) (any, error)
	GetAscendants(models.GetAscendants) (any, error)
	GetAscendantsAndDescendants(models.GetAscendantsAndDescendants) (any, error)
	GetAscendantsAndChildren(models.GetAscendantsAndChildren) (any, error)

	PostPerson(models.PostPerson) (any, error)
	PostParentRelationship(models.PostParentRelationship) (any, error)

	DelPerson(models.DelPerson) (any, error)
	DelParentRelationship(models.DelParentRelationship) (any, error)
}
