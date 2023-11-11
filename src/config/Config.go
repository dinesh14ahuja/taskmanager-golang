package config

import (
	"log"

	"github.com/spf13/viper"
)

var config *viper.Viper

var AppConfig *viper.Viper

func Init() {
	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName("config")
	config.AddConfigPath("src/config")
	err := config.ReadInConfig()
	if err != nil {
		log.Fatalln("Error on parsing configuration file", err.Error())
	}
}

func GetConfig() *viper.Viper {
	return config
}
