//go:build ignore

package rest

import (
	"github.com/gin-gonic/gin"

	"genealogy-tree/internal/debug"
	"genealogy-tree/internal/models"
)

func deserializePrototype(ctx *gin.Context) (models.Prototype, error) {
	var (
		err     error
		payload models.Prototype
	)

	if err = ctx.Bind(&payload); err != nil {
		debug.ShowErr("deserializePrototype", "Failed to deserialize request payload", err)
		return models.Prototype{}, err
	}

	return payload, nil
}
