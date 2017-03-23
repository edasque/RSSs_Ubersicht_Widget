RSSs-Widget
===========

Übersicht Stock Widget for displaying multiple RSS feeds

* Can override header for RSS feed through config.json (name property)
* Can override max number of items per feed or globally in config (default is 5)
* Adjust left and top in the style section in index.coffee if you want it placed somewhere else on the screen
* Customize the feed list by editing RSSs-Widget/RSSs.widget/config.json
* JSON config file should be pretty self-explanatory:

```JSON
{
	"max_items": 5,
	"feeds": [
	{
		"name": "Boston Globe",
		"url": "http://www.bostonglobe.com/rss/feedly.xml",
		"max_items": 5},
	{
			"name": "BBC Top News",
		"url": "http://feeds.bbci.co.uk/news/rss.xml?edition=int"
	}, {
		"name": "The Verge",
		"url": "http://www.theverge.com/rss/frontpage"
	},
	{
		"name": "WSJ World News",
		"url": "http://www.wsj.com/xml/rss/3_7085.xml",
				"max_items": 8
	}]
}
```