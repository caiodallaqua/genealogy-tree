package main

import (
	"genealogy-tree/env"
	"genealogy-tree/internal/adapters/api"
	"genealogy-tree/internal/adapters/database"
	"genealogy-tree/internal/adapters/rest"
	"genealogy-tree/internal/ports"
	"log"
)

var config *env.AppConfig

func init() {
	config = env.NewConfig()
}

// @title        Genealogy Tree API
// @version      0.1
// @description  API for handling person entity and its relationships in a genealogy tree.

// @host      localhost:8080
// @BasePath  /

func main() {
	// Adapters must obey their ports' interfaces.
	// This allows the internal code to depend mainly
	// on abstractions rather than concrete implementations.
	var (
		dbAdapter   ports.DBPort
		apiAdapter  ports.APIPort
		restAdapter ports.RestPort
	)

	var err error

	dbAdapter, err = database.NewAdapter(config.DB.Addr, config.DB.User, config.DB.Pwd)
	if err != nil {
		log.Printf("| main | Failed to connect to database: %v", err)
		return
	}

	apiAdapter = api.NewAdapter(dbAdapter)
	restAdapter = rest.NewAdapter(config.Addr, apiAdapter)

	router := restAdapter.CreateRouter()
	if err = restAdapter.RunRestServer(router); err != nil {
		log.Printf("| main | Failed to launch rest server: %v", err)
	}
}
