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
	CallGetStatus() error
	CallPostPerson(models.PostPerson)
	CallPostParentRelationship(models.PostParentRelationship)
	CallDelParentRelationship(models.DelParentRelationship)

	// GET
	CallGetPerson(models.GetPerson)
	CallGetAscendants(models.GetAscendants)
}

// Driven

type DBPort interface {
	GetStatus() error
	PostPerson(models.PostPerson)
	PostParentRelationship(models.PostParentRelationship)
	DelParentRelationship(models.DelParentRelationship)

	// GET
	GetPerson(models.GetPerson)
	GetAscendants(models.GetAscendants)
}
