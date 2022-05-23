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
	CallTest()
	CallAddPerson(models.AddPerson)
	CallAddParentRelationship(models.AddParentRelationship)
	CallGetAscendants(models.GetAscendants)
}

// Driven

type DBPort interface {
	Test()
	AddPerson(models.AddPerson)
	AddParentRelationship(models.AddParentRelationship)
	GetAscendants(models.GetAscendants)
}
