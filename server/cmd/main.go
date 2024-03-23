package main

import (
	"log"

	config "github.com/R894/bugtracker"
	"github.com/R894/bugtracker/internal/database"
	"github.com/R894/bugtracker/internal/web"
	"github.com/R894/bugtracker/storage"
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
