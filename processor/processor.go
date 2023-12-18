package processor

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
	"match_statistics_scrapper/db"
	"match_statistics_scrapper/db/repos"
	"match_statistics_scrapper/models"
	"match_statistics_scrapper/notifier"
	"match_statistics_scrapper/scrapper"
	"match_statistics_scrapper/utils"
	"strings"
)

func FetchPlayerStats() {
	ps := repos.PlayersStore{DB: db.ConnectDB()}
	ss := repos.StatsStore{DB: db.ConnectDB()}

	playersData, err := ps.GetPlayersData()
	if err != nil {
		log.Println("Failed to get players data")
		return
	}

	for _, player := range playersData {
		url := player.Url
		dataScrapped := scrapFromUrl(url)

		for _, data := range dataScrapped {
			data.PlayerName = player.PlayerName
			data.Source = url

			stringified := utils.StringifyStruct(*data)
			hashed := utils.GetHashOfData(stringified)

			statsData, err := ss.FindHash(hashed)
			if errors.Is(err, gorm.ErrRecordNotFound) && statsData == nil {
				err = notifier.TgBot(stringified)
				if err != nil {
					fmt.Println("Failed while sending the update")
					continue
				}
				err = ss.InsertData(&models.ScrappedData{
					Hash: hashed,
					Data: stringified,
				})
				if err != nil {
					return
				}
			}
			if err != nil {
				return
			}
		}
	}
}

func scrapFromUrl(url string) []*models.MatchStatResponse {
	if strings.Contains(url, "basketball.eurobasket.com") {
		//return scrapper.ScrapsEuroBasket(url)
	} else if strings.Contains(url, "espn.co.uk") {
		return scrapper.EspnScrap(url)
	} else if strings.Contains(url, "nbl.com.au") {
		//return scrapper.NblScrap(url)
	} else if strings.Contains(url, "bleague.jp") {
		return scrapper.ScrapsBleague(url)
	} else if strings.Contains(url, "bnxtleague.com") {
		return scrapper.ScrapsBnxt(url)
	} else if strings.Contains(url, "britishbasketballleague.com") {
		return scrapper.ScrapsBritishBasketBall(url)
	} else if strings.Contains(url, "210.140.77.209") {
		return scrapper.ScrapsB3league(url)
	}

	fmt.Println("URL not supported !")
	return nil
}
