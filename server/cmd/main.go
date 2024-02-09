package main

import (
	"bugtracker/internal/database"
	"bugtracker/internal/web"
	"bugtracker/storage"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	db, err := database.NewPostgresDB()
	if err != nil {
		log.Fatal("Error initializing database: ", err)
	}
	storage.Migrate(db.Db)

	err = web.StartServer(db.Db)
	if err != nil {
		log.Fatal("Error starting web server: ", err)
	}
}
