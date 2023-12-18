package scrapper

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"match_statistics_scrapper/models"
	"match_statistics_scrapper/utils"
	"strings"
)

func ScrapsBleague(url string) []*models.MatchStatResponse {
	// creating a new Colly instance
	c := colly.NewCollector()

	var rows [][]string

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: ", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Page visited: ", r.Request.URL)
	})

	// scraping logic section.gamelogWidget
	c.OnHTML("#scores", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, row *colly.HTMLElement) {
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

	stats := []*models.MatchStatResponse{}
	for _, row := range rows {
		statResp := &models.MatchStatResponse{
			Date:   utils.BleagueDate(strings.TrimSpace(row[0])),
			Opp:    strings.TrimSpace(row[1]),
			Result: strings.TrimSpace(row[3]),
			Min:    strings.TrimSpace(row[5]),
			FGP:    strings.TrimSpace(row[9]),
			FTP:    strings.TrimSpace(row[18]),
			ThreeP: strings.TrimSpace(row[15]),
			REB:    strings.TrimSpace(row[23]),
			AST:    strings.TrimSpace(row[24]),
			BLK:    strings.TrimSpace(row[28]),
			STL:    strings.TrimSpace(row[27]),
			PF:     strings.TrimSpace(row[31]),
			TO:     strings.TrimSpace(row[26]),
			PTS:    strings.TrimSpace(row[6]),
		}
		stats = append(stats, statResp)
	}
	fmt.Println("Bleague Scrapper result : ", stats)
	return stats
}
