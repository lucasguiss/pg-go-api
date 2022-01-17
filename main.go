package main

import (
	"github.com/lucasguiss/pg-go-api/database"
	"github.com/lucasguiss/pg-go-api/server"
)

func main() {
	database.StartDB()

	server := server.NewServer()
	server.Run()
}
