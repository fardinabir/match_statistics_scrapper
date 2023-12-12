package main

import (
	"github.com/spf13/viper"
	"log"
	"match_statistics_scrapper/job"
)

func main() {
	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	job.StartTicker()
}
