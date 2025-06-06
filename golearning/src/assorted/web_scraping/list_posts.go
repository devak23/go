package web_scraping

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
)

func postScrape() {
	doc, err := goquery.NewDocument("http://jonathanmh.com")
	if err != nil {
		log.Fatal(err)
	}

	// use CSS selector found with the browser inspector
	doc.Find("#main article .entry-title").Each(func(index int, item *goquery.Selection) {
		title := item.Text()
		linkTag := item.Find("a")
		link, _ := linkTag.Attr("href")
		fmt.Printf("Post #%d: %s - %s\n", index, title, link)
	})
}

func ListPostsMain() {
	postScrape()
}
