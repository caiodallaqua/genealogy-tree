package rest

import (
	"log"

	"github.com/gin-gonic/gin"

	"genealogy-tree/internal/models"
)

func deserializeNewPerson(ctx *gin.Context) (models.NewPerson, error) {
	var err error
	var payload models.NewPerson

	if err = ctx.Bind(&payload); err != nil {
		log.Println(err)
		return models.NewPerson{}, err
	}

	return payload, nil
}
