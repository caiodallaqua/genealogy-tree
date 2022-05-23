package rest

import (
	"log"

	"github.com/gin-gonic/gin"

	"genealogy-tree/internal/models"
)

func deserializeGetAscendants(ctx *gin.Context) (models.GetAscendants, error) {
	var err error
	var payload models.GetAscendants

	if err = ctx.Bind(&payload); err != nil {
		log.Println(err)
		return models.GetAscendants{}, err
	}

	return payload, nil
}

func deserializeGetPerson(ctx *gin.Context) (models.GetPerson, error) {
	var err error
	var payload models.GetPerson

	if err = ctx.Bind(&payload); err != nil {
		log.Println(err)
		return models.GetPerson{}, err
	}

	return payload, nil
}

func deserializePostPerson(ctx *gin.Context) (models.PostPerson, error) {
	var err error
	var payload models.PostPerson

	if err = ctx.Bind(&payload); err != nil {
		log.Println(err)
		return models.PostPerson{}, err
	}

	return payload, nil
}

func deserializePostParentRelationship(ctx *gin.Context) (models.PostParentRelationship, error) {
	var err error
	var payload models.PostParentRelationship

	if err = ctx.Bind(&payload); err != nil {
		log.Println(err)
		return models.PostParentRelationship{}, err
	}

	return payload, nil
}

func deserializeDelParentRelationship(ctx *gin.Context) (models.DelParentRelationship, error) {
	var err error
	var payload models.DelParentRelationship

	if err = ctx.Bind(&payload); err != nil {
		log.Println(err)
		return models.DelParentRelationship{}, err
	}

	return payload, nil
}
