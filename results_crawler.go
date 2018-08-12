package main

import (
	"net/http"
	"github.com/golang/glog"
	"github.com/PuerkitoBio/goquery"
	"fmt"
)

type ResultsCrawler struct{ resultsPageUrls []string }

func (crawler ResultsCrawler) LoadResults() ([]string) {
	ch := make(chan []string)
	e := make(chan error)
	urlSet := make(map[string]bool)

	for i, resultsPageUrl := range crawler.resultsPageUrls {
		fmt.Printf("Loading results %d/%d\n", i+1, len(crawler.resultsPageUrls))
		go crawler.loadResultsForEvent(resultsPageUrl, ch, e)
		select {
		case profileUrls := <-ch:
			for _, profileUrl := range profileUrls {
				urlSet[profileUrl] = true
			}
		case _ = <-e:
		}
	}

	uniqueProfileUrls := make([]string, 0, len(urlSet))
	for k := range urlSet {
		uniqueProfileUrls = append(uniqueProfileUrls, k)
	}

	return uniqueProfileUrls
}

func (crawler ResultsCrawler) loadResultsForEvent(url string, ch chan []string, e chan error){
	// Make HTTP request
	response, err := http.Get(url)
	if err != nil {
		glog.Error("Failed to load event results", err)
	}
	defer response.Body.Close()

	// Create a goquery document from the HTTP response
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		glog.Fatal("Error parsing event results page. ", err)
		e <- err
	}

	// Find all the competitors
	coupleLinks := document.Find("a.coupleLink")

	links := make([]string, coupleLinks.Length(), coupleLinks.Length())
	coupleLinks.Each(func(i int, element *goquery.Selection) {
		id, _ := element.Attr("encname")
		links[i] = crawler.coupleUrl(id)
	})

	ch <- links
}

func (crawler ResultsCrawler) coupleUrl(id string)(string){
	return fmt.Sprintf("https://www.dancesportinfo.net/Couple/%s/Details", id)
}
