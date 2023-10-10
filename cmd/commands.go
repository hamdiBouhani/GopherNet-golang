package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/hamdiBouhani/GopherNet-golang/server"
	"github.com/hamdiBouhani/GopherNet-golang/services"
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

			go func() {
				d := 10 * time.Minute
				err := burrowServiceInstance.Report(d)
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
