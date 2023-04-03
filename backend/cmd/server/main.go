package main

import (
	"log"
	"net/http"

	"myapp/api/routes"
)

func main() {
	e := routes.New()

	s := http.Server{
		Addr:    ":8080",
		Handler: e,
	}

	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

// /tepmlates/list  {[]}

//
//// {
//	"name": "name",
//	"language": "language" // main language
//	"description": ""
//	"create_at": ""
//	"update_at": ""
////}
