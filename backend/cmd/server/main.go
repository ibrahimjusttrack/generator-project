package main

import (
	"log"
	"myapp/internal/db"
	"net/http"

	"myapp/internal/routes"
)

func main() {
	e := routes.New()
	db.InitDb()

	s := http.Server{
		Addr:    ":8080",
		Handler: e,
	}

	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
