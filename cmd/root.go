package cmd

import (
	"log"

	"gshop/internal/config"
	"gshop/internal/gorm"
	"gshop/sdk"
	"gshop/sdk/logger"

	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "Gshop",
	Short: "Gshop API",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := config.Init(); err != nil {
			log.Fatalf("%s", err.Error())

			return err
		}

		logger.InitServLogger(true)

		db, err := gorm.InitDb()

		if err != nil {
			log.Fatalf("%s", err.Error())

			return err
		}

		sc := sdk.New(db)

		server := NewServer(sc)

		if err = server.Run(); err != nil {
			log.Fatalf("%s", err.Error())

			return err
		}

		return nil
	},
}

func Execute() {
	rootCmd.AddCommand(migrateCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
