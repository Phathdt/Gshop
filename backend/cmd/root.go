package cmd

import (
	"os"

	"context"
	"log"

	"github.com/spf13/cobra"
	"gshop/internal/application/server"

	"gshop/internal/application/services"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "Gshop",
	Short: "Gshop",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		sc, err := services.NewServiceContext(ctx)
		if err != nil {
			log.Fatalf("%s", err.Error())
		}

		s := server.NewServer(sc)

		if err = s.Run(); err != nil {
			log.Fatalf("%s", err.Error())
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
