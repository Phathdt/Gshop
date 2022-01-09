package cmd

import (
	"log"

	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
	"gshop/internal/config"
	"gshop/internal/gorm"
	"gshop/sdk"
)

var rootCmd = &cobra.Command{
	Use:   "Gshop",
	Short: "Gshop API",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := config.Init(); err != nil {
			log.Fatalf("%s", err.Error())

			return err
		}

		db, err := gorm.InitDb()

		if err != nil {
			log.Fatalf("%s", err.Error())

			return err
		}

		sc := sdk.NewServiceContext(db)

		server := NewServer(sc)

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
	}
}
