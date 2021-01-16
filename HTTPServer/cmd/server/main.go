package main

import (
	"github.com/kjbrock/proglog/HTTPServer/internal/server"
	"github.com/kjbrock/proglog/HTTPServer/internal/log"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
