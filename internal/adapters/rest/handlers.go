package rest

import (
	"genealogy-tree/internal/debug"
	"genealogy-tree/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// -------------------------------------     GET     -------------------------------------

// Ping DB to health check (maybe)
func (restAdapter Adapter) getStatus(ctx *gin.Context) {
	var (
		err        error
		resPayload any
	)

	if resPayload, err = restAdapter.api.CallGetStatus(); err != nil {
		debug.ShowErr("getStatus", "Failed to get status from DB", err)
		response(ctx, http.StatusInternalServerError, ErrInternal)

		return
	}

	response(ctx, http.StatusOK, resPayload)
}

// ShowAccount godoc
// @Summary      Get Person
// @Description  get person by ID
// @Tags         Person
// @Param        id query uint true "Person ID"
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.GetPersonRes
// @Failure      422  {object}  rest.GenTreeError
// @Failure      500  {object}  rest.GenTreeError
// @Failure      400  {object}  rest.GenTreeError
// @Router       /person [get]
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

// ShowAccount godoc
// @Summary      Get Ascendants
// @Description  get all ascendants by person ID
// @Tags         Ascendants / Descendants
// @Param        id query int true "Person ID"
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.GetAscendantsRes
// @Failure      422  {object}  rest.GenTreeError
// @Failure      500  {object}  rest.GenTreeError
// @Router       /ascendants [get]
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

// ShowAccount godoc
// @Summary      Get Ascendants and Descendants
// @Description  get all ascendants and descendants by person ID
// @Tags         Ascendants / Descendants
// @Param        id query int true "Person ID"
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.GetAscendantsAndDescendantsRes
// @Failure      422  {object}  rest.GenTreeError
// @Failure      500  {object}  rest.GenTreeError
// @Router       /ascendants-and-descendants [get]
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

// ShowAccount godoc
// @Summary      Get Ascendants and Immediate Children
// @Description  give person ID. get all ascendants and children of the leaf ascendants
// @Tags         Ascendants / Descendants
// @Param        id query int true "Person ID"
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.GetAscendantsAndChildrenRes
// @Failure      422  {object}  rest.GenTreeError
// @Failure      500  {object}  rest.GenTreeError
// @Router       /ascendants-and-children [get]
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

// ShowAccount godoc
// @Summary      Post Person
// @Description  add person given name and optionally birth
// @Tags         Person
// @Param        PostPerson body models.PostPerson true "Person data"
// @Accept       json
// @Produce      json
// @Success      201  {object}  models.PostPersonRes
// @Failure      422  {object}  rest.GenTreeError
// @Failure      500  {object}  rest.GenTreeError
// @Router       /person [post]
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

// ShowAccount godoc
// @Summary      Post Parent Relationship
// @Description  add relationship given two ids (of parent and child)
// @Tags         Relation
// @Param        ParentRelationship body models.PostParentRelationship true "Parent relationship data"
// @Accept       json
// @Produce      json
// @Success      201  {object}  models.PostParentRelationshipRes
// @Failure      422  {object}  rest.GenTreeError
// @Failure      500  {object}  rest.GenTreeError
// @Router       /parent-relationship [post]
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

// ShowAccount godoc
// @Summary      Delete Person
// @Description  removes person and its relationships given the person id
// @Tags         Person
// @Param        DelPerson body models.DelPerson true "Removes person from database"
// @Accept       json
// @Produce      json
// @Success      201  {object}  models.DelPersonRes
// @Failure      422  {object}  rest.GenTreeError
// @Failure      500  {object}  rest.GenTreeError
// @Router       /person [delete]
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

// ShowAccount godoc
// @Summary      Delete Parent Relationship
// @Description  removes relationship between parent and child given their ids
// @Tags         Relation
// @Param        ParentRelationship body models.DelParentRelationship true "Parent relationship"
// @Accept       json
// @Produce      json
// @Success      201  {object}  models.DelParentRelationshipRes
// @Failure      422  {object}  rest.GenTreeError
// @Failure      500  {object}  rest.GenTreeError
// @Router       /parent-relationship [delete]
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
