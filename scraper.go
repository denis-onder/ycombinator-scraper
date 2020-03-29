package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func find(_ int, element *goquery.Selection) {
	titleElement := element.Find(".title > a.storylink")
	index := element.Find(".rank").Text()
	title := titleElement.Text()
	link, _ := titleElement.Attr("href")
	fmt.Println(index, title, "\n", link, "\n")
}

func scrape() (error, bool) {
	const url = "https://news.ycombinator.com/"
	res, err := http.Get(url)
	if err != nil {
		return err, false
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		return err, false
	}

	doc.Find(".athing").Each(find)

	return nil, true
}
