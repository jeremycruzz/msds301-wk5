package wikiscrape

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/gocolly/colly"
)

type Scraper struct {
	colly  *colly.Collector
	corpus []Data
}

// create a new scraper with a concurrency limit
func NewScraper(concurrency int) *Scraper {

	c := colly.NewCollector(
		colly.Async(true),
		colly.MaxDepth(0))

	c.Limit(&colly.LimitRule{
		Parallelism: concurrency,
	})

	scraper := &Scraper{colly: c}
	scraper.init()

	return scraper
}

// make callback for colly
func (s *Scraper) init() {
	s.colly.OnHTML("title", func(element *colly.HTMLElement) {
		fmt.Println("starting")
		title := element.Text
		text := element.DOM.Closest("body").Text()

		// get tags from url
		url := element.Request.URL.String()
		tags := strings.Split(url, "/")

		data := Data{
			Url:   url,
			Title: title,
			Text:  text,
			Tags:  tags[1:],
		}
		fmt.Println("stopping")
		// append to the corpus
		s.corpus = append(s.corpus, data)
	})
}

// scrapes url(s)
func (s *Scraper) Scrape(urls ...string) {
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go s.scrapeURL(url, &wg)
	}
	wg.Wait()
}

// writes the corpus to a file
func (s *Scraper) WriteCorpusToFile(filepath string) error {

	// marshall json with indents for readability
	corpusJSON, err := json.MarshalIndent(s.corpus, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling to JSON: %v", err)
	}

	// write to file
	err = os.WriteFile(filepath, corpusJSON, 0644)
	if err != nil {
		return fmt.Errorf("error writing JSON to file: %v", err)
	}

	return nil
}

func (s *Scraper) scrapeURL(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	err := s.colly.Visit(url)
	if err != nil {
		fmt.Printf("Error visiting %s: %v\n", url, err)
	}
}
