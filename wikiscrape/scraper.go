package wikiscrape

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/gocolly/colly"
	"github.com/jeremycruzz/msds301-wk5/sets"
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
	s.colly.OnHTML("html", func(element *colly.HTMLElement) {
		// extract the title from head
		title := element.ChildText("head title")

		// extract the text from body
		text := element.ChildText("body")

		// get tags from url
		url := element.Request.URL.String()
		tags := extractTags(url)

		data := Data{
			Url:   url,
			Title: title,
			Tags:  tags,
			Text:  text,
		}

		// append to the corpus
		s.corpus = append(s.corpus, data)
	})
}

// scrapes url(s)
func (s *Scraper) Scrape(urls ...string) {

	for _, url := range urls {
		s.colly.Visit(url)
	}
	s.colly.Wait() // never forget this
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

func extractTags(url string) []string {
	var cleanedTags []string
	var cleanedLastTags []string

	// remove parenthesis
	url = strings.ReplaceAll(url, "(", "")
	url = strings.ReplaceAll(url, ")", "")

	// split by slashes
	tags := strings.Split(url, "/")

	// filter out empty tags
	for _, tag := range tags {
		if tag != "" {
			cleanedTags = append(cleanedTags, tag)
		}
	}

	// split the last tag by underscores
	lastTag := cleanedTags[len(cleanedTags)-1]
	lastTagSplit := strings.Split(lastTag, "_")

	// remove stop words and empty tags
	for _, tag := range lastTagSplit {
		if tag != "" && !sets.StopWords[strings.ToLower(tag)] {
			cleanedLastTags = append(cleanedLastTags, strings.ToLower(tag))
		}
	}

	// replace last part with split
	cleanedTags = append(cleanedTags[:len(cleanedTags)-2], cleanedLastTags...)

	return cleanedTags
}
