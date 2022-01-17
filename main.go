package main

import "github.com/lucasguiss/pg-go-api/server"

func main() {
	server := server.NewServer()
	server.Run()
}
