package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hamdiBouhani/GopherNet-golang/cmd"
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

	if err := cmd.CommandRoot(burrowServiceInstance).Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(2)
	}

}
