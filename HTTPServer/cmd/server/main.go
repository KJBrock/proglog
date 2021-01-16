package main

import (
	"log"
	"github.com/kjbrock/proglog/HTTTPServer/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
