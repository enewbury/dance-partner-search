package main

import (
	"sync"
	"os"
	"log"
	"bufio"
	"fmt"
)

func LoadPartnershipData() (map[string][]*Partnership){
	//getCachedPartnerships()
	ch := make(chan *CrawlResult)
	var visited sync.Map

	coupleLinks := ResultsCrawler{resultsPageUrls: getEvents()}.LoadResults()
	for _, link := range coupleLinks {
		go ProfileCrawler{ch: ch}.LoadProfile(link, visited)
	}

	followerMap := make(map[string][]*Partnership)

	numPartnerships := len(coupleLinks)
	for i := 0; i < numPartnerships; i++ {
		result := <- ch
		if result.newPartnership {
			numPartnerships++
		}
		followerMap[result.partnership.follower] = append(followerMap[result.partnership.follower], result.partnership)
		fmt.Printf("Loaded %d/%d couple profile\n", i+1, numPartnerships)
	}

	return followerMap
}

func getEvents() ([]string){
	eventUrls := make([]string, 0)

	file, err := os.Open("events.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		eventUrls = append(eventUrls, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return eventUrls
}
