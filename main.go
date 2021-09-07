package main

import (
	"log"

	"github.com/provinite/oh-ya/server"
)

func main() {
	log.Print("Starting server on http://localhost:8080")
	log.Fatal(server.StartServer())
}
