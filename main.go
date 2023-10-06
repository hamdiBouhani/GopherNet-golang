package main

import (
	"log"

	"github.com/hamdiBouhani/GopherNet-golang/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = server.NewHttpService().Start()
	if err != nil {
		log.Fatalf(err.Error())
	}

}
