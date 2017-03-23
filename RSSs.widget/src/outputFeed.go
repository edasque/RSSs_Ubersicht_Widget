package main

import (
	"encoding/json"
	"fmt"
	"github.com/mmcdole/gofeed"
	"io/ioutil"
	"os"
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
	feed, _ := fp.ParseURL(theFeed.URL)

	if theFeed.Name != "" {
		fmt.Printf("<H3>%s</H3>\n", theFeed.Name)

	} else {

		fmt.Printf("<H3>%s</H3>\n", feed.Title)

	}

	fmt.Println("<ul>")

	items := feed.Items[:max_items]

	for _, element := range items {
		fmt.Printf("<li>%s - <a href='%s'> <i class='fa fa-external-link'></i> </a></li>\n", element.Title, element.Link)

	}
	fmt.Println("</ul>")
}
