package api

import (
	"genealogy-tree/internal/models"
	"genealogy-tree/internal/ports"
)

type Adapter struct {
	db ports.DBPort
}

func NewAdapter(db ports.DBPort) Adapter {
	return Adapter{
		db: db,
	}
}

// -------------------------------------     GET     -------------------------------------

func (apiAdapter Adapter) CallGetStatus() error {
	return apiAdapter.db.GetStatus()
}

func (apiAdapter Adapter) CallGetPerson(payload models.GetPerson) (any, error) {
	return apiAdapter.db.GetPerson(payload)
}

func (apiAdapter Adapter) CallGetAscendants(payload models.GetAscendants) (any, error) {
	return apiAdapter.db.GetAscendants(payload)
}

func (apiAdapter Adapter) CallGetAscendantsAndDescendants(payload models.GetAscendantsAndDescendants) (any, error) {
	return apiAdapter.db.GetAscendantsAndDescendants(payload)
}

func (apiAdapter Adapter) CallGetAscendantsAndChildren(payload models.GetAscendantsAndChildren) (any, error) {
	return apiAdapter.db.GetAscendantsAndChildren(payload)
}

// -------------------------------------     POST     -------------------------------------

func (apiAdapter Adapter) CallPostPerson(payload models.PostPerson) (any, error) {
	return apiAdapter.db.PostPerson(payload)
}

func (apiAdapter Adapter) CallPostParentRelationship(payload models.PostParentRelationship) (any, error) {
	return apiAdapter.db.PostParentRelationship(payload)
}

// -------------------------------------     DELETE     -------------------------------------

func (apiAdapter Adapter) CallDelPerson(payload models.DelPerson) (any, error) {
	return apiAdapter.db.DelPerson(payload)
}

func (apiAdapter Adapter) CallDelParentRelationship(payload models.DelParentRelationship) (any, error) {
	return apiAdapter.db.DelParentRelationship(payload)
}
