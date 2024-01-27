package web

import (
	"bugtracker/internal/web/router"
	"fmt"
	"net/http"
	"os"
)

func StartServer() {
	port := os.Getenv("API_PORT")
	if port == "" {
		port = "3333"
	}

	r := router.SetupRouter()
	address := fmt.Sprintf(":%s", port)

	http.ListenAndServe(address, r)
}
