package rest

import (
	"genealogy-tree/internal/debug"
	"genealogy-tree/internal/ports"

	"github.com/clbanning/mxj/v2"
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
	router.GET("/ascendants-and-descendants", restAdapter.getAscendantsAndDescendantsHandler)
	router.GET("/ascendants-and-children", restAdapter.getAscendantsAndChildrenHandler)

	// POST
	router.POST("/person", restAdapter.postPersonHandler)
	router.POST("/parent-relationship", restAdapter.postParentRelationshipHandler)

	// DELETE
	router.DELETE("/person", restAdapter.delPersonHandler)
	router.DELETE("/parent-relationship", restAdapter.delParentRelationshipHandler)

	return router
}

func (restAdapter Adapter) RunRestServer(router *gin.Engine) error {
	var err error = router.Run(restAdapter.addr)

	if err != nil {
		debug.ShowErr("RunRestServer", "Failed to run rest server", err)
	}

	return err
}

func response(ctx *gin.Context, statusCode int, resPayload any) {
	var contentType string = ctx.Request.Header.Get("Content-Type")

	switch contentType {
	case "application/xml":
		xml, err := mxj.AnyXmlIndent(resPayload, "", "  ", "genealogy-tree")
		if err != nil {
			debug.ShowErr("response", "Failed to serialize response to XML", err)
			return
		}
		ctx.Data(statusCode, contentType, xml)
	case "application/yaml":
		ctx.YAML(statusCode, resPayload)
	default:
		ctx.JSON(statusCode, resPayload)
	}
}
