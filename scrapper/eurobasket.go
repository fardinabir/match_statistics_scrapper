package scrapper

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"match_statistics_scrapper/models"
	"match_statistics_scrapper/utils"
)

func ScrapsEuroBasket(url string) {
	// creating a new Colly instance
	c := colly.NewCollector()

	var rows [][]string

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: ", err)
	})

	// scraping logic section.gamelogWidget
	c.OnHTML("tbody", func(e *colly.HTMLElement) {
		e.ForEach("#23", func(_ int, row *colly.HTMLElement) {
			var cells []string

			row.ForEach("td", func(_ int, cell *colly.HTMLElement) {
				cells = append(cells, cell.Text)
			})

			rows = append(rows, cells)
		})
	})

	// visiting the target page
	err := c.Visit(url)
	if err != nil {
		fmt.Println("Couldn't Visit")
	}

	stats := []models.EuroBasketStat{}
	for _, row := range rows {
		data := models.EuroBasketStat{
			Date:        utils.EuroBasketDate(row[0]),
			Team:        row[1],
			AgainstTeam: row[2],
			Result:      row[3],
			Min:         row[4],
			Pts:         row[5],
			TwoFGP:      row[6],
			ThreeFGP:    row[7],
			FT:          row[8],
			RO:          row[9],
			RD:          row[10],
			RT:          row[11],
			AS:          row[12],
			PF:          row[13],
			BS:          row[14],
			ST:          row[15],
			TO:          row[16],
			RNK:         row[17],
		}
		stats = append(stats, data)
	}
	fmt.Println("EuroBasketStat Scrapper result : ", stats)
	return
}
