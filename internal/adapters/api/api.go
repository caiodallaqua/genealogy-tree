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

func (apiAdapter Adapter) CallTest() {
	apiAdapter.db.Test()
}

func (apiAdapter Adapter) CallAddPerson(payload models.AddPerson) {
	apiAdapter.db.AddPerson(payload)
}

func (apiAdapter Adapter) CallAddParentRelationship(payload models.AddParentRelationship) {
	apiAdapter.db.AddParentRelationship(payload)
}

func (apiAdapter Adapter) CallGetAscendants(payload models.GetAscendants) {
	apiAdapter.db.GetAscendants(payload)
}
