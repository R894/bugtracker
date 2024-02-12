package main

import (
	config "bugtracker"
	"bugtracker/internal/database"
	"bugtracker/internal/web"
	"bugtracker/storage"
	"log"
)

func main() {
	config.FetchEnvVars()
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
