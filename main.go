package main

import (
	"net/http"

	"github.com/Stance0102/gin-api/internal/routers"
)

func main() {
	router := routers.NewRouter()

	s := &http.Server{
		Addr:    ":80",
		Handler: router,
	}

	s.ListenAndServe()
}
