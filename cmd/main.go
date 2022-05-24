package main

import (
	"genealogy-tree/internal/adapters/api"
	"genealogy-tree/internal/adapters/database"
	"genealogy-tree/internal/adapters/rest"
	"genealogy-tree/internal/ports"

	"genealogy-tree/env"
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
	//docs.SwaggerInfo.Title = "Genealogy Tree API"
	//docs.SwaggerInfo.Host = "localhost:8998"
	//docs.SwaggerInfo.BasePath = "/docs"
	//docs.SwaggerInfo.Version = "0.1"

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
		return
	}

	apiAdapter = api.NewAdapter(dbAdapter)
	restAdapter = rest.NewAdapter(config.Addr, apiAdapter)

	router := restAdapter.CreateRouter()
	restAdapter.RunRestServer(router) // see this later, error return
}
