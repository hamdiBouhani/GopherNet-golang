package cmd

import (
	"os"

	"github.com/hamdiBouhani/GopherNet-golang/services"
	"github.com/spf13/cobra"
)

func CommandRoot(burrowServiceInstance *services.BurrowService) *cobra.Command {
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
