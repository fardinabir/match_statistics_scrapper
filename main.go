package main

import (
	"match_statistics_scrapper/scrapper"
)

func main() {
	//viper.SetConfigFile("config.yaml")
	//if err := viper.ReadInConfig(); err != nil {
	//	log.Fatalf("Error reading config file: %s", err)
	//}
	//
	//job.StartTicker()
	scrapper.Scraps()
}
