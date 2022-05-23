package rest

import (
	"log"

	"github.com/gin-gonic/gin"
)

// Ping DB to health check (maybe)
func (restAdapter Adapter) status(ctx *gin.Context) {
	log.Println("42")
}

func (restAdapter Adapter) addPersonHandler(ctx *gin.Context) {
	payload, err := deserializeAddPerson(ctx)
	if err != nil {
		return
	}

	log.Println(payload)

	restAdapter.api.CallAddPerson(payload)
}

func (restAdapter Adapter) addParentRelationshipHandler(ctx *gin.Context) {
	payload, err := deserializeAddParentRelationship(ctx)
	if err != nil {
		return
	}

	log.Println(payload)

	restAdapter.api.CallAddParentRelationship(payload)
}

func (restAdapter Adapter) getAscendantsHandler(ctx *gin.Context) {
	payload, err := deserializeGetAscendants(ctx)
	if err != nil {
		return
	}

	log.Println(payload)

	restAdapter.api.CallGetAscendants(payload)
}
