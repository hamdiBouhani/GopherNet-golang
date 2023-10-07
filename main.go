package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/hamdiBouhani/GopherNet-golang/server"
	"github.com/hamdiBouhani/GopherNet-golang/services"
	"github.com/hamdiBouhani/GopherNet-golang/storage/pg"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

func NewRestCmd(burrowServiceInstance *services.BurrowService) *cobra.Command {
	return &cobra.Command{
		Use:   "rest",
		Short: "Run rest server",
		Long:  ``,
		Run: func(commandServe *cobra.Command, args []string) {

			go func() {

				err := burrowServiceInstance.RunUpdateStatusTask(1 * time.Minute)
				if err != nil {
					log.Fatalf(err.Error())
				}

			}()

			err := server.NewHttpService(burrowServiceInstance).Start()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(2)
			}
		},
	}
}

func NewLoadBurrowsCmd(burrowServiceInstance *services.BurrowService) *cobra.Command {
	return &cobra.Command{
		Use:   "load-burrows",
		Short: "Load Burrows from initial.json",
		Long:  ``,
		Run: func(commandServe *cobra.Command, args []string) {

			err := burrowServiceInstance.InitialBurrowStates()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(2)
			}
		},
	}
}

func commandRoot(burrowServiceInstance *services.BurrowService) *cobra.Command {
	rootCmd := &cobra.Command{
		Use: "gopherne",
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
			os.Exit(2)
		},
	}

	rootCmd.AddCommand(NewRestCmd(burrowServiceInstance))
	rootCmd.AddCommand(NewLoadBurrowsCmd(burrowServiceInstance))

	return rootCmd
}

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

	if err := commandRoot(burrowServiceInstance).Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(2)
	}

}
