package main

import (
	"encoding/json"
	"fmt"
	"github.com/mmcdole/gofeed"
	"io/ioutil"
	"os"
)

func main() {
	file, e := ioutil.ReadFile("./config.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	type feed struct {
		name string `json:"name"`
	}

	type config struct {
		feeds []feed `json:"feeds"`
	}

	localConfig := config{}
	json.Unmarshal(file, &config)

	outputAndParseFeed("http://feeds.bbci.co.uk/news/rss.xml?edition=int")

}

func outputAndParseFeed(feedURL string) {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(feedURL)
	fmt.Printf("<H3>%s</H3>\n", feed.Title)
	fmt.Println("<ul>")

	for index, element := range feed.Items {
		if index < 6 {
			fmt.Printf("<li>%s - <a href='%s'> <i class='fa fa-external-link'></i> </a></li>\n", element.Title, element.Link)

		}
	}
	fmt.Println("</ul>")
}
