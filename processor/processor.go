package processor

import (
	"fmt"
	"match_statistics_scrapper/notifier"
	"match_statistics_scrapper/scrapper"
)

func FetchPlayerStats() {
	dataScrapped := scrapper.EspnScrap("https://www.espn.co.uk/mens-college-basketball/player/_/id/5105785/max-mackinnon")

	for _, data := range dataScrapped {
		err := notifier.TgBot(*data)
		if err != nil {
			fmt.Println("Failed while sending the update")
			continue
		}

	}
}
