package main

import "match_statistics_scrapper/scrapper"

func main() {
	//viper.SetConfigFile("config.yaml")
	//if err := viper.ReadInConfig(); err != nil {
	//	log.Fatalf("Error reading config file: %s", err)
	//}
	//
	//job.StartTicker()
	//scrapper.ScrapsBnxt(`https://bnxtleague.com/en/player-statistics/?player_id=2882&amp;team\_id=162`)
	//scrapper.EspnScrap("https://www.espn.co.uk/mens-college-basketball/player/_/id/5105785/max-mackinnon")
	//scrapper.ScrapsEuroBasket("https://basketball.eurobasket.com/player/Jacob-Lampkin/335999")
	scrapper.NblScrap("https://nbl.com.au/player/3713/853140/lachlan-anderson")
}
