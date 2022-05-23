package rest

import (
	"log"

	"github.com/gin-gonic/gin"

	"genealogy-tree/internal/models"
)

func deserializeAddPerson(ctx *gin.Context) (models.AddPerson, error) {
	var err error
	var payload models.AddPerson

	if err = ctx.Bind(&payload); err != nil {
		log.Println(err)
		return models.AddPerson{}, err
	}

	return payload, nil
}

func deserializeAddParentRelationship(ctx *gin.Context) (models.AddParentRelationship, error) {
	var err error
	var payload models.AddParentRelationship

	if err = ctx.Bind(&payload); err != nil {
		log.Println(err)
		return models.AddParentRelationship{}, err
	}

	return payload, nil
}

func deserializeGetAscendants(ctx *gin.Context) (models.GetAscendants, error) {
	var err error
	var payload models.GetAscendants

	if err = ctx.Bind(&payload); err != nil {
		log.Println(err)
		return models.GetAscendants{}, err
	}

	return payload, nil
}
