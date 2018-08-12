package main

import (
	"net/http"
	"github.com/golang/glog"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"strconv"
	"time"
	"sync"
)

type ProfileCrawler struct {
	ch chan * CrawlResult
	document *goquery.Document
}


func (crawler ProfileCrawler) LoadProfile(profileUrl string, visited sync.Map){
	visited.Store(profileUrl, true)

	response, err := http.Get(profileUrl)
	if err != nil {
		glog.Error("Failed to load profile results: ", err)
	}
	defer response.Body.Close()

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		glog.Error("Error parsing partnership page: ", err)
	}

	crawler.document = document

	partnership := Partnership {
		follower:		 crawler.getFollowerName(),
		names:           crawler.getNames(),
		highScore:       crawler.getBallroomScore(),
		lastCompetition: crawler.getLastCompetition(),
		profileUrl:      profileUrl,
	}

	newPartnershipUrl := crawler.getNewPartnershipUrl()
	_, found := visited.Load(newPartnershipUrl)
	if newPartnershipUrl != "" && !found {
		go ProfileCrawler{ch: crawler.ch}.LoadProfile(newPartnershipUrl, visited)
	}

	crawler.ch <- &CrawlResult{partnership: &partnership, newPartnership: newPartnershipUrl != "" && !found}
}

func (crawler ProfileCrawler) getNames() string {
	return crawler.document.Find("h2").First().Children().Text()
}

func (crawler ProfileCrawler) getFollowerName() string {
	return crawler.document.Find("#ContentPlaceHolder1_HerName").First().Text()
}

func (crawler ProfileCrawler) getBallroomScore() int {
	span := crawler.document.Find("#ContentPlaceHolder1_rpRating_lblStyle_0")
	if strings.TrimSpace(span.Text()) == "Ballroom" {
		p := span.Parent().Parent()
		scoreString := p.Children().Remove().End().Text()
		scoreString = strings.Replace(scoreString, "-", "", -1)
		scoreString = strings.TrimSpace(scoreString)
		score, _ := strconv.Atoi(scoreString)
		return score
	} else {
		return 0
	}
}

func (crawler ProfileCrawler) getLastCompetition() time.Time{
	compDateString := crawler.document.Find("#ContentPlaceHolder1_LastComp").Find("p2").Text()
	timeOfLastComp, err := time.Parse("Monday, 02 January 2006", compDateString)
	if err != nil {
		glog.Error("Unable to parse date:", err)
	}
	return timeOfLastComp
}

func (crawler ProfileCrawler) getNewPartnershipUrl() string {
	url, _ := crawler.document.Find("#ContentPlaceHolder1_currLadyLink").First().Attr("href")
	if url != "" {
		return "https://www.dancesportinfo.net" + url
	} else {
		return ""
	}
}
