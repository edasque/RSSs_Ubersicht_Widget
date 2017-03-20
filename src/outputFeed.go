package main

import (
	"encoding/json"
	"fmt"
	"github.com/mmcdole/gofeed"
	"io/ioutil"
	"os"
)

type feed struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type config struct {
	Feeds []feed `json:"feeds"`
}

func main() {
	file, e := ioutil.ReadFile("RSSs.widget/config.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	localConfig := config{}
	if err := json.Unmarshal(file, &localConfig); err != nil {
		panic(err)
	}

	// fmt.Println(localConfig)

	for _, item := range localConfig.Feeds {
		outputAndParseFeed(item)

	}

}

func outputAndParseFeed(theFeed feed) {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(theFeed.URL)
	fmt.Printf("<H3>%s</H3>\n", theFeed.Name)
	fmt.Println("<ul>")

	for index, element := range feed.Items {
		if index < 6 {
			fmt.Printf("<li>%s - <a href='%s'> <i class='fa fa-external-link'></i> </a></li>\n", element.Title, element.Link)

		}
	}
	fmt.Println("</ul>")
}
