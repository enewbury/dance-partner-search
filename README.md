# partner-search
This is a practice project in Go to find potential dance partners from dancesportinfo results.  Given a list of event urls, the program loads all the couples entered, finds any new couples that have formed since that competition, and returns a curated list of partners who have not competed in the desired number of months, or who are competing with a partner at a significantly lower ranking than they competed with before.

## Specify Events
Include a file called events.txt in the directory you are running the application from.  This file should contain a list of urls, one on each line, for the results pages of dancesportinfo.net of your desired events to crawl

## Run
```bash
go build
./partner_search
```

## Chagne filtering
Edit the main.go file to change the number of monthsSinceLastCompetition, and the number of points different from a previous partnership in order to affect how many results are returned.
