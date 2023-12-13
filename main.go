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
	scrapper.EspnScrap("https://www.espn.co.uk/mens-college-basketball/player/_/id/5105785/max-mackinnon")
	scrapper.ScrapsEuroBasket("https://basketball.eurobasket.com/player/Jacob-Lampkin/335999")
}
