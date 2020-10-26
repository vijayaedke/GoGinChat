package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func ReadConfig() error{
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$GOPATH/goginchat/config")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			return viper.ConfigFileNotFoundError{}
		} else {
			// Config file was found but another error was produced
			return fmt.Errorf("Error parsing file")
		}
	}
	return nil
}