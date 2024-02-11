package main

import (
	utils "bugtracker/internal"
	"bugtracker/internal/database"
	"bugtracker/internal/web"
	"bugtracker/storage"
	"log"
)

func main() {
	utils.FetchEnvVars()
	postgres, err := database.NewPostgresDB()
	if err != nil {
		log.Fatal("Error initializing database: ", err)
	}
	storage.Migrate(postgres.Db)

	err = web.StartServer(postgres.Db)
	if err != nil {
		log.Fatal("Error starting web server: ", err)
	}
}
