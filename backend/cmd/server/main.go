package main

import (
	"log"
	"net/http"

	"myapp/api/routes"
	"myapp/db"
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
