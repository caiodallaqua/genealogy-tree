package rest

import (
	"log"

	"github.com/gin-gonic/gin"
)

// Ping DB to health check (maybe)
func (restAdapter Adapter) status(ctx *gin.Context) {
	log.Println("42")
}

func (restAdapter Adapter) newPerson(ctx *gin.Context) {
	payload, err := deserializeNewPerson(ctx)
	if err != nil {
		return
	}

	log.Println(payload)
	restAdapter.api.CallTest()
}
