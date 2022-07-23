package rest

import (
	"github.com/gin-gonic/gin"

	"genealogy-tree/internal/debug"
	"genealogy-tree/internal/models"
)

func deserializeGetPerson(ctx *gin.Context) (models.GetPerson, error) {
	var (
		err     error
		payload models.GetPerson
	)

	if err = ctx.Bind(&payload); err != nil {
		debug.ShowErr("deserializeGetPerson", "Failed to deserialize request payload", err)
		return models.GetPerson{}, err
	}

	return payload, nil
}

func deserializeGetAscendants(ctx *gin.Context) (models.GetAscendants, error) {
	var (
		err     error
		payload models.GetAscendants
	)

	if err = ctx.Bind(&payload); err != nil {
		debug.ShowErr("deserializeGetAscendants", "Failed to deserialize request payload", err)
		return models.GetAscendants{}, err
	}

	return payload, nil
}

func deserializeGetAscendantsAndDescendants(ctx *gin.Context) (models.GetAscendantsAndDescendants, error) {
	var (
		err     error
		payload models.GetAscendantsAndDescendants
	)

	if err = ctx.Bind(&payload); err != nil {
		debug.ShowErr("deserializeGetAscendantsAndDescendants", "Failed to deserialize request payload", err)
		return models.GetAscendantsAndDescendants{}, err
	}

	return payload, nil
}

func deserializeGetAscendantsAndChildren(ctx *gin.Context) (models.GetAscendantsAndChildren, error) {
	var (
		err     error
		payload models.GetAscendantsAndChildren
	)

	if err = ctx.Bind(&payload); err != nil {
		debug.ShowErr("deserializeGetAscendantsAndChildren", "Failed to deserialize request payload", err)
		return models.GetAscendantsAndChildren{}, err
	}

	return payload, nil
}

func deserializePostPerson(ctx *gin.Context) (models.PostPerson, error) {
	var (
		err     error
		payload models.PostPerson
	)

	if err = ctx.Bind(&payload); err != nil {
		debug.ShowErr("deserializePostPerson", "Failed to deserialize request payload", err)
		return models.PostPerson{}, err
	}

	return payload, nil
}

func deserializePostParentRelationship(ctx *gin.Context) (models.PostParentRelationship, error) {
	var (
		err     error
		payload models.PostParentRelationship
	)

	if err = ctx.Bind(&payload); err != nil {
		debug.ShowErr("deserializePostParentRelationship", "Failed to deserialize request payload", err)
		return models.PostParentRelationship{}, err
	}

	return payload, nil
}

func deserializeDelPerson(ctx *gin.Context) (models.DelPerson, error) {
	var (
		err     error
		payload models.DelPerson
	)

	if err = ctx.Bind(&payload); err != nil {
		debug.ShowErr("deserializeDelPerson", "Failed to deserialize request payload", err)
		return models.DelPerson{}, err
	}

	return payload, nil
}

func deserializeDelParentRelationship(ctx *gin.Context) (models.DelParentRelationship, error) {
	var (
		err     error
		payload models.DelParentRelationship
	)

	if err = ctx.Bind(&payload); err != nil {
		debug.ShowErr("deserializeDelParentRelationship", "Failed to deserialize request payload", err)
		return models.DelParentRelationship{}, err
	}

	return payload, nil
}
