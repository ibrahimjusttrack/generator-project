package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	s := http.Server{
    Addr:        ":8080",
    Handler:     e,
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
