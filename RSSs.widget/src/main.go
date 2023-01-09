package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"

	"github.com/mmcdole/gofeed"
)

type feed struct {
	Name      string `json:"name"`
	URL       string `json:"url"`
	Max_Items int    `json:"max_items"`
}

type config struct {
	Feeds     []feed `json:"feeds"`
	Max_Items int    `json:"max_items"`
}

func main() {
	file, e := ioutil.ReadFile("RSSs-Widget/RSSs.widget/config.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	localConfig := config{}
	if err := json.Unmarshal(file, &localConfig); err != nil {
		panic(err)
	}

	// fmt.Println(localConfig)

	const DEFAULT_MAX = 5

	var max = DEFAULT_MAX

	if localConfig.Max_Items > 0 {
		max = localConfig.Max_Items
	}

	for _, item := range localConfig.Feeds {

		if item.Max_Items > 0 {
			outputAndParseFeed(item, item.Max_Items)
		} else {

			outputAndParseFeed(item, max)
		}

	}

}

func outputAndParseFeed(theFeed feed, max_items int) {
	fp := gofeed.NewParser()
	feed, feedParseError := fp.ParseURL(theFeed.URL)

	if feedParseError != nil {
		fmt.Println("Error querying: " + theFeed.URL)
		log.Fatalln(feedParseError)

	}

	if theFeed.Name != "" {
		fmt.Printf("<H3>%s</H3>\n", theFeed.Name)

	} else {

		fmt.Printf("<H3>%s</H3>\n", feed.Title)

	}

	feedLength := len(feed.Items)

	fmt.Println("<ul>")

	var number_number_of_item = int(math.Min(float64(feedLength), float64(max_items)))

	items := feed.Items[:number_number_of_item]

	for _, element := range items {
		fmt.Printf("<li><span>%s</span> - <a href='%s'> <i class='fa fa-external-link'></i> </a></li>\n", element.Title, element.Link)

	}
	fmt.Println("</ul>")
}
