package main

import (
	"log"

	"github.com/kjbrock/proglog/HTTPServer/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
