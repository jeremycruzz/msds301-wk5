package main

import (
	"fmt"
	"time"

	"github.com/jeremycruzz/msds301-wk5/wikiscrape"
)

var URLs = []string{
	"https://en.wikipedia.org/wiki/Robotics",
	"https://en.wikipedia.org/wiki/Robot",
	"https://en.wikipedia.org/wiki/Reinforcement_learning",
	"https://en.wikipedia.org/wiki/Robot_Operating_System",
	"https://en.wikipedia.org/wiki/Intelligent_agent",
	"https://en.wikipedia.org/wiki/Software_agent",
	"https://en.wikipedia.org/wiki/Robotic_process_automation",
	"https://en.wikipedia.org/wiki/Chatbot",
	"https://en.wikipedia.org/wiki/Applications_of_artificial_intelligence",
	"https://en.wikipedia.org/wiki/Android_(robot)",
}

func main() {

	startTime := time.Now()

	//create scraper
	concurrency := 4
	scraper := wikiscrape.NewScraper(concurrency)
	outputFilePath := "./results/corpus.json"

	scraper.Scrape(URLs...)

	err := scraper.WriteCorpusToFile(outputFilePath)
	if err != nil {
		fmt.Printf("Error writing corpus to file: %v\n", err)
	} else {
		fmt.Printf("Corpus data has been saved to %s\n", outputFilePath)
	}

	totalTime := time.Since(startTime).Nanoseconds()
	fmt.Printf("Scraping completed in %d ns\n", totalTime)
}
