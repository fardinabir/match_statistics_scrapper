package scrapper

import (
	"fmt"
	"github.com/go-rod/rod"
	"github.com/gocolly/colly"
	"log"
)

func ScrapsBnxt() {
	// creating a new Colly instance
	url := `https://bnxtleague.com/en/player-statistics/?player_id=2882&amp;team\_id=162`

	page := rod.New().MustConnect().MustPage(url).MustWaitLoad()

	// Get the HTML content after JavaScript execution
	page.MustHTML()

	page.MustClose()

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
	c.OnHTML("body", func(e *colly.HTMLElement) {
		println(e)
		e.ForEach("tr", func(_ int, row *colly.HTMLElement) {
			var cells []string

			row.ForEach("td", func(_ int, cell *colly.HTMLElement) {
				cells = append(cells, cell.Text)
			})

			rows = append(rows, cells)
		})
	})

	// visiting the target page

	println(page.String())

	err := c.Visit(url)
	if err != nil {
		fmt.Println("Couldn't Visit")
	}

	fmt.Println(rows)
	fmt.Println("---------------------------------------------")
	fmt.Println("---------------------------------------------")
	for _, row := range rows {
		fmt.Println(row)
	}
	return
}
