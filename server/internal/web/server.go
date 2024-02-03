package web

import (
	"bugtracker/internal/web/router"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
)

func StartServer(db *sql.DB) error {
	port := os.Getenv("API_PORT")
	if port == "" {
		port = "3333"
	}

	r, err := router.SetupRouter(db)
	if err != nil {
		return err
	}
	address := fmt.Sprintf(":%s", port)
	log.Println("Starting server on address ", address)
	if err := http.ListenAndServe(address, r); err != nil {
		return err
	}
	return nil
}
