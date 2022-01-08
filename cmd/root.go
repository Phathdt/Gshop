package cmd

import (
	"gshop/internal/config"
	"gshop/internal/postgresql"
	"gshop/sdk"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gshop",
	Short: "A brief description of your application",

	RunE: func(cmd *cobra.Command, args []string) error {
		if err := config.Init(); err != nil {
			log.Fatalf("%s", err.Error())

			return err
		}

		db, err := postgresql.InitDb()

		if err != nil {
			log.Fatalf("%s", err.Error())

			return err
		}

		sc := sdk.ServiceConfig{
			DB: db,
		}

		server := NewServer(&sc)

		if err = server.Run(); err != nil {
			log.Fatalf("%s", err.Error())

			return err
		}

		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)

		os.Exit(1)
	}
}
