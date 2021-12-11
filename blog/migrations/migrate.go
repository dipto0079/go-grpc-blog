package main

import (
	"grpc-todos/todo/storage/postgres"
	"log"
)

func main() {
	if err := postgres.Migrate(); err != nil {
		log.Fatal(err)
	}
}
