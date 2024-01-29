package main

import (
	"bugtracker/internal/database"
	"bugtracker/internal/web"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := database.NewPostgresDB()
	if err != nil {
		log.Fatal(err)
	}
	// storage.Migrate(db.Db)

	err = web.StartServer(db)
	if err != nil {
		log.Fatal(err)
	}
}
