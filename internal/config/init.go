package config

import "github.com/spf13/viper"

func Init() error {
	viper.SetConfigFile("app.yaml")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	viper.AutomaticEnv()

	return nil
}
