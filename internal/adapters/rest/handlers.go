package rest

import (
	"log"

	"github.com/gin-gonic/gin"
)

// -------------------------------------     GET     -------------------------------------

// Ping DB to health check (maybe)
func (restAdapter Adapter) getStatus(ctx *gin.Context) {
	err := restAdapter.api.CallGetStatus()

	if err != nil {
		return
	}

	log.Println("42")
}

func (restAdapter Adapter) getPersonHandler(ctx *gin.Context) {
	payload, err := deserializeGetPerson(ctx)
	if err != nil {
		return
	}

	restAdapter.api.CallGetPerson(payload)
}

func (restAdapter Adapter) getAscendantsHandler(ctx *gin.Context) {
	payload, err := deserializeGetAscendants(ctx)
	if err != nil {
		return
	}

	restAdapter.api.CallGetAscendants(payload)
}

// -------------------------------------     POST     -------------------------------------

func (restAdapter Adapter) postPersonHandler(ctx *gin.Context) {
	payload, err := deserializePostPerson(ctx)
	if err != nil {
		return
	}

	log.Println(payload)

	restAdapter.api.CallPostPerson(payload)
}

func (restAdapter Adapter) postParentRelationshipHandler(ctx *gin.Context) {
	payload, err := deserializePostParentRelationship(ctx)
	if err != nil {
		return
	}

	log.Println(payload)

	restAdapter.api.CallPostParentRelationship(payload)
}

// -------------------------------------     DELETE     -------------------------------------

func (restAdapter Adapter) delParentRelationshipHandler(ctx *gin.Context) {
	payload, err := deserializeDelParentRelationship(ctx)
	if err != nil {
		return
	}

	log.Println(payload)

	restAdapter.api.CallDelParentRelationship(payload)
}
