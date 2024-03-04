package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
    "path/filepath"
	"html/template"
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
	pwd, _ := os.Executable()
	file, e := ioutil.ReadFile(filepath.Dir(pwd)+"/../config.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	localConfig := config{}
	if err := json.Unmarshal(file, &localConfig); err != nil {
		panic(err)
	}

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

func noescape(str string) template.HTML {
return template.HTML(str)
}

func outputAndParseFeed(theFeed feed, max_items int) {
	fp := gofeed.NewParser()
	pwd, _ := os.Executable()
	feed, feedParseError := fp.ParseURL(theFeed.URL)

	if feedParseError != nil {
		fmt.Println("Error querying: " + theFeed.URL)
		log.Fatalln(feedParseError)

	}

	if theFeed.Name != "" {
		feed.Title= theFeed.Name
	} 
	
	Tmpl, err := template.New("").Funcs(template.FuncMap{"noescape":noescape }).ParseFiles(filepath.Dir(pwd)+"/../feed.tmpl")
    
	if err != nil {
        log.Fatal("Feed parsing error:", err)
    }


	fmt.Println(Tmpl.ExecuteTemplate(os.Stdout, "feed.tmpl", feed))
	
}
