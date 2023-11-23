package main

import (
	"handmedown-backend/src/config"
	"handmedown-backend/src/routes"
)

func init() {
	var err error
	config.DB, err = config.InitializeDB()
	if err != nil {
		panic(err)
	}

	config.MigrateDB(config.DB)
}

func main() {
	defer config.DisconnectDB(config.DB)

	r := routes.SetRoutes(config.DB)
	r.Run()
}
