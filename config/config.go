package config

import (
	"log"

	"github.com/Jsagudelo1704/Go/structs"
	"github.com/spf13/viper"
)

func GetConfig() structs.Configuration {
	conf := structs.Configuration{}

	viper.SetConfigFile("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = viper.Unmarshal(&conf)
	if err != nil {
		log.Fatal(err)
	}

	return conf
}
