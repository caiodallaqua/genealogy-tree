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

func (apiAdapter Adapter) CallGetAscendants(payload models.GetAscendants) {
	apiAdapter.db.GetAscendants(payload)
}

func (apiAdapter Adapter) CallGetPerson(payload models.GetPerson) {
	apiAdapter.db.GetPerson(payload)
}

// -------------------------------------     POST     -------------------------------------

func (apiAdapter Adapter) CallPostPerson(payload models.PostPerson) {
	apiAdapter.db.PostPerson(payload)
}

func (apiAdapter Adapter) CallPostParentRelationship(payload models.PostParentRelationship) {
	apiAdapter.db.PostParentRelationship(payload)
}

// -------------------------------------     DELETE     -------------------------------------

func (apiAdapter Adapter) CallDelParentRelationship(payload models.DelParentRelationship) {
	apiAdapter.db.DelParentRelationship(payload)
}
