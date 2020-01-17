package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var googleDomains = map[string]string{
	"com": "https://www.google.com",
	"uk":  "https://www.google.co.uk",
	"ru":  "https://www.google.ru",
	"fr":  "https://www.google.fr",
}

var searchPrefix = "/search?q="

func main() {
	searchTerm, domain, countryCode, languageCode, debug := cliParameters()
	googleResults, err := GoogleScrape(searchTerm, countryCode, languageCode, debug)
	if err == nil {
		writeCSVData(googleResults, domain+".csv")
	} else {
		fmt.Println("Error occured", err)
	}
}

func buildGoogleURL(searchTerm string, countryCode string, languageCode string) string {
	searchTerm = strings.Trim(searchTerm, " ")
	searchTerm = strings.Replace(searchTerm, " ", "+", -1)

	googleBase := googleDomains["com"] + searchPrefix
	if googleDomain, found := googleDomains[countryCode]; found {
		googleBase = googleDomain + searchPrefix
	}

	return fmt.Sprintf("%s%s&num=100&hl=%s", googleBase, searchTerm, languageCode)
}

func buildGoogleURLPageLink(pageLink string, countryCode string) string {
	googleBase := googleDomains["com"]
	if googleDomain, found := googleDomains[countryCode]; found {
		googleBase = googleDomain
	}
	fmt.Println(googleBase)
	return fmt.Sprintf("%s%s", googleBase, pageLink)
}

func googleRequest(searchURL string) (*http.Response, error) {
	baseClient := &http.Client{}
	req, _ := http.NewRequest("GET", searchURL, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36")
	res, err := baseClient.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func googleResultParser(response *http.Response) ([]GoogleResult, []string, error) {
	doc, err := goquery.NewDocumentFromResponse(response)
	if err != nil {
		return nil, nil, err
	}
	results := parseResultsFromPage(doc)
	pagesLinks := []string{}
	linksPages := doc.Find("a.fl")
	for i := range linksPages.Nodes {
		item := linksPages.Eq(i)
		pageLink, _ := item.Attr("href")
		pagesLinks = append(pagesLinks, pageLink)
	}
	return results, pagesLinks, err
}

func parseResultsFromPage(doc *goquery.Document) []GoogleResult {
	results := []GoogleResult{}
	sel := doc.Find("div.g")
	for i := range sel.Nodes {
		item := sel.Eq(i)
		linkTag := item.Find("a")
		link, _ := linkTag.Attr("href")
		titleTag := item.Find("h3")
		descTag := item.Find("span.st")
		desc := descTag.Text()
		title := titleTag.Text()
		link = strings.Trim(link, " ")
		if link != "" && link != "#" {
			result := GoogleResult{
				link,
				title,
				desc,
			}
			results = append(results, result)
		}
	}
	return results
}

// GoogleScrape scrapes data from Google search engine
func GoogleScrape(searchTerm string, countryCode string, languageCode string, debug bool) ([]GoogleResult, error) {
	googleURL := buildGoogleURL(searchTerm, countryCode, languageCode)
	if debug {
		fmt.Println(googleURL)
	}
	res, err := googleRequest(googleURL)
	if err != nil {
		return nil, err
	}
	scrapes, pagesLinks, err := googleResultParser(res)
	for _, pl := range pagesLinks {
		if strings.HasPrefix(pl, searchPrefix) {
			randomSleep(5)
			googleURL = buildGoogleURLPageLink(pl, countryCode)
			fmt.Printf("Adding additional links from the page: %s\n", googleURL)
			res, err := googleRequest(googleURL)
			if err != nil {
				return nil, err
			}
			scrapesOtherPages, _, err := googleResultParser(res)
			scrapes = append(scrapes, scrapesOtherPages...)
			if err != nil {
				return nil, err
			}
		}
	}
	if err != nil {
		return nil, err
	}
	return scrapes, nil
}

func randomSleep(maxSleepTime int) {
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Duration(rand.Intn(maxSleepTime)) * time.Second)
}
