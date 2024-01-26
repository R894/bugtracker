package web

import (
	"bugtracker/internal/web/router"
	"net/http"
)

func StartServer() {
	r := router.SetupRouter()
	http.ListenAndServe(":3333", r)
}
