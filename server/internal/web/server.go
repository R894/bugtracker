package web

import (
	"bugtracker/internal/database"
	"bugtracker/internal/web/router"
	"fmt"
	"net/http"
	"os"
)

func StartServer(db *database.PostgresDB) error {
	port := os.Getenv("API_PORT")
	if port == "" {
		port = "3333"
	}

	r, err := router.SetupRouter(db)
	if err != nil {
		return err
	}
	address := fmt.Sprintf(":%s", port)

	if err := http.ListenAndServe(address, r); err != nil {
		return err
	}

	return nil
}
