package cmd

import (
	"log"

	"gshop/config"
	"gshop/pkg"
	"gshop/pkg/gorm"
	"gshop/pkg/logger"
	"gshop/pkg/redis"

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

		logger.New()

		db, err := gorm.InitDb()

		if err != nil {
			log.Fatalf("%s", err.Error())

			return err
		}

		rdb, err := redis.NewRedis()
		if err != nil {
			log.Fatalf("%s", err.Error())
			return err
		}

		sc := pkg.New(db, rdb)

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
