package main

import (
	"log"

	"github.com/hamdiBouhani/GopherNet-golang/server"
	"github.com/hamdiBouhani/GopherNet-golang/services"
	"github.com/hamdiBouhani/GopherNet-golang/storage/pg"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := pg.NewDBConn()
	err = db.CreateConnection()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = db.Migrate()
	if err != nil {
		log.Fatalf(err.Error())
	}

	burrowServiceInstance := services.NewBurrowService(db)

	err = server.NewHttpService(burrowServiceInstance).Start()
	if err != nil {
		log.Fatalf(err.Error())
	}

}
