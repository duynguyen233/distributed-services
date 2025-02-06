package main

import (
	"log"

	"github.com/duynguyen233/proglog/chapter1/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
