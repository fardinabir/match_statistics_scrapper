package scrapper

import (
	"encoding/csv"
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"os"
)

func Scraps() {
	// creating a new Colly instance
	c := colly.NewCollector()

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
	c.OnHTML("Table__TBODY", func(e *colly.HTMLElement) {
		fmt.Println(e)
		fmt.Println("0---------------0")
		fmt.Println(e.ChildText("18"))
	})

	// visiting the target page
	c.Visit("https://www.espn.co.uk/mens-college-basketball/player/_/id/5105785/max-mackinnon")
	return
}

// initializing a data structure to keep the scraped data
type PokemonProduct struct {
	url, image, name, price string
}

func Scrap() {
	// initializing the slice of structs to store the data to scrape
	var pokemonProducts []PokemonProduct

	// creating a new Colly instance
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: ", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Page visited: ", r.Request.URL)
	})

	// scraping logic
	c.OnHTML("li.product", func(e *colly.HTMLElement) {
		pokemonProduct := PokemonProduct{}

		pokemonProduct.url = e.ChildAttr("a", "href")
		pokemonProduct.image = e.ChildAttr("img", "src")
		pokemonProduct.name = e.ChildText("h2")
		pokemonProduct.price = e.ChildText(".price")

		pokemonProducts = append(pokemonProducts, pokemonProduct)
	})

	// visiting the target page
	c.Visit("https://scrapeme.live/shop/")

	// opening the CSV file
	file, err := os.Create("products.csv")
	if err != nil {
		log.Fatalln("Failed to create output CSV file", err)
	}
	defer file.Close()

	// initializing a file writer
	writer := csv.NewWriter(file)

	// writing the CSV headers
	headers := []string{
		"url",
		"image",
		"name",
		"price",
	}
	writer.Write(headers)

	// writing each Pokemon product as a CSV row
	for _, pokemonProduct := range pokemonProducts {
		// converting a PokemonProduct to an array of strings
		record := []string{
			pokemonProduct.url,
			pokemonProduct.image,
			pokemonProduct.name,
			pokemonProduct.price,
		}

		// adding a CSV record to the output file
		writer.Write(record)
	}
	defer writer.Flush()
	return
}
