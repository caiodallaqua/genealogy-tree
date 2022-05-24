package rest

import (
	"genealogy-tree/internal/models"
	"log"
	"net/http"

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
	var (
		err error

		// Request / Response content
		reqPayload models.GetPerson
		resPayload any
	)

	if reqPayload, err = deserializeGetPerson(ctx); err != nil {
		response(ctx, http.StatusUnprocessableEntity, ErrFailedToDeserialize)
		return
	}

	if resPayload, err = restAdapter.api.CallGetPerson(reqPayload); err != nil {
		response(ctx, http.StatusInternalServerError, ErrInternal)
		return
	}

	if resPayload == nil {
		response(ctx, http.StatusBadRequest, ErrResourceNotFound)
		return
	}

	response(ctx, http.StatusOK, resPayload)
}

func (restAdapter Adapter) getAscendantsHandler(ctx *gin.Context) {
	var (
		err error

		// Request / Response content
		reqPayload models.GetAscendants
		resPayload any
	)

	if reqPayload, err = deserializeGetAscendants(ctx); err != nil {
		response(ctx, http.StatusUnprocessableEntity, ErrFailedToDeserialize)
		return
	}

	if resPayload, err = restAdapter.api.CallGetAscendants(reqPayload); err != nil {
		response(ctx, http.StatusInternalServerError, ErrInternal)
		return
	}

	response(ctx, http.StatusOK, resPayload)
}

func (restAdapter Adapter) getAscendantsAndDescendantsHandler(ctx *gin.Context) {
	var (
		err error

		// Request / Response content
		reqPayload models.GetAscendantsAndDescendants
		resPayload any
	)

	if reqPayload, err = deserializeGetAscendantsAndDescendants(ctx); err != nil {
		response(ctx, http.StatusUnprocessableEntity, ErrFailedToDeserialize)
		return
	}

	if resPayload, err = restAdapter.api.CallGetAscendantsAndDescendants(reqPayload); err != nil {
		response(ctx, http.StatusInternalServerError, ErrInternal)
		return
	}

	response(ctx, http.StatusOK, resPayload)
}

func (restAdapter Adapter) getAscendantsAndChildrenHandler(ctx *gin.Context) {
	var (
		err error

		// Request / Response content
		reqPayload models.GetAscendantsAndChildren
		resPayload any
	)

	if reqPayload, err = deserializeGetAscendantsAndChildren(ctx); err != nil {
		response(ctx, http.StatusUnprocessableEntity, ErrFailedToDeserialize)
		return
	}

	if resPayload, err = restAdapter.api.CallGetAscendantsAndChildren(reqPayload); err != nil {
		response(ctx, http.StatusInternalServerError, ErrInternal)
		return
	}

	response(ctx, http.StatusOK, resPayload)
}

// -------------------------------------     POST     -------------------------------------

func (restAdapter Adapter) postPersonHandler(ctx *gin.Context) {
	var (
		err error

		// Request / Response content
		reqPayload models.PostPerson
		resPayload any
	)

	if reqPayload, err = deserializePostPerson(ctx); err != nil {
		response(ctx, http.StatusUnprocessableEntity, ErrFailedToDeserialize)
		return
	}

	if resPayload, err = restAdapter.api.CallPostPerson(reqPayload); err != nil {
		response(ctx, http.StatusInternalServerError, ErrInternal)
		return
	}

	response(ctx, http.StatusCreated, resPayload)
}

func (restAdapter Adapter) postParentRelationshipHandler(ctx *gin.Context) {
	var (
		err error

		// Request / Response content
		reqPayload models.PostParentRelationship
		resPayload any
	)

	if reqPayload, err = deserializePostParentRelationship(ctx); err != nil {
		response(ctx, http.StatusUnprocessableEntity, ErrFailedToDeserialize)
		return
	}

	if resPayload, err = restAdapter.api.CallPostParentRelationship(reqPayload); err != nil {
		response(ctx, http.StatusInternalServerError, ErrInternal)
		return
	}

	response(ctx, http.StatusCreated, resPayload)
}

// -------------------------------------     DELETE     -------------------------------------

func (restAdapter Adapter) delPersonHandler(ctx *gin.Context) {
	var (
		err error

		// Request / Response content
		reqPayload models.DelPerson
		resPayload any
	)

	if reqPayload, err = deserializeDelPerson(ctx); err != nil {
		response(ctx, http.StatusUnprocessableEntity, ErrFailedToDeserialize)
		return
	}

	if resPayload, err = restAdapter.api.CallDelPerson(reqPayload); err != nil {
		response(ctx, http.StatusInternalServerError, ErrInternal)
		return
	}

	response(ctx, http.StatusCreated, resPayload)
}

func (restAdapter Adapter) delParentRelationshipHandler(ctx *gin.Context) {
	var (
		err error

		// Request / Response content
		reqPayload models.DelParentRelationship
		resPayload any
	)

	if reqPayload, err = deserializeDelParentRelationship(ctx); err != nil {
		response(ctx, http.StatusUnprocessableEntity, ErrFailedToDeserialize)
		return
	}

	if resPayload, err = restAdapter.api.CallDelParentRelationship(reqPayload); err != nil {
		response(ctx, http.StatusInternalServerError, ErrInternal)
		return
	}

	response(ctx, http.StatusCreated, resPayload)
}
