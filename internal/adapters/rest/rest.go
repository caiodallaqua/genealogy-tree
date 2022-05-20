package rest

import (
	"genealogy-tree/internal/ports"

	"github.com/gin-gonic/gin"
)

type Adapter struct {
	addr string
	api  ports.APIPort
}

func NewAdapter(addr string, api ports.APIPort) Adapter {
	return Adapter{
		addr: addr,
		api:  api,
	}
}

func (restAdapter Adapter) CreateRouter() *gin.Engine {
	router := gin.Default()

	// GET
	router.GET("/status", restAdapter.status)

	// POST
	router.POST("/person", restAdapter.newPerson)

	return router
}

func (restAdapter Adapter) RunRestServer(router *gin.Engine) error {
	var err error = router.Run(restAdapter.addr)

	return err
}