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
	router.GET("/status", restAdapter.getStatus)
	router.GET("/person", restAdapter.getPersonHandler)
	router.GET("/ascendants", restAdapter.getAscendantsHandler)

	// POST
	router.POST("/person", restAdapter.postPersonHandler)
	router.POST("/parent-relationship", restAdapter.postParentRelationshipHandler)

	// DELETE
	router.DELETE("/parent-relationship", restAdapter.delParentRelationshipHandler)

	return router
}

func (restAdapter Adapter) RunRestServer(router *gin.Engine) error {
	var err error = router.Run(restAdapter.addr)

	return err
}
