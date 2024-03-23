package web

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/R894/bugtracker/internal/web/router"
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
