package api

import "genealogy-tree/internal/ports"

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
