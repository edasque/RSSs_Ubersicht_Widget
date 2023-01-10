# RSSs-Widget for Übersicht

[Übersicht](http://tracesof.net/uebersicht/) Widget for displaying multiple RSS feeds on the desktop. 

Still using  the [classic CoffeeScript widget format](https://github.com/felixhageloh/uebersicht/blob/master/ClassicWidgets.md), as the AppleScript support is still broken on the React / JSX widgets.

---

"Übersicht lets you run system commands and display their output on your Mac desktop in little containers, called widgets."

---

![Alt text](screenshot.png?raw=true "Preview")

## Installation

1. Copy the RSSs.widget folder to the **Übersicht widget folder**.

2. Override the heading for RSS feeds through **config.json** (name property).

3. Override max number of items per feed or globally in **config.json** (default is 5).

4. Change styling and positioning  in **index.coffee**. 

5. Customize the feeds list by editing **config.json**

6. Customize list rendering through **list.tmpl**
   
   ```
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
    }]
   }
   ```

## Render Template

feed.tmpl

```<section>
 <h3>{{ .Title}}</h3>
 <ul>
 {{range .Items}}
  <li>
   <span>{{ .Title}}</span><a href='{{ .Link}}'>
   <i class='fa fa-solid fa-link'></i> </a>
  </li>
 {{end}}
 </ul>
</section>
```

### Available [gofeed](https://github.com/mmcdole/gofeed) Tags

Those can be used within the {{range .Items}} section.

| `gofeed.Item` | RSS                                                                                                                                                                                 | Atom                                                                          | JSON                                 |
| ------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ------------------------------------ |
| Title         | /rss/channel/item/title<br>/rdf: RDF/item/title<br>/rdf: RDF/item/dc:title<br>/rss/channel/item/dc:title                                                                            | /feed/entry/title                                                             | /items/title                         |
| Description   | /rss/channel/item/description<br>/rdf: RDF/item/description<br>/rss/channel/item/dc:description<br>/rdf: RDF/item/dc:description                                                    | /feed/entry/summary                                                           | /items/summary                       |
| Content       | /rss/channel/item/content:encoded                                                                                                                                                   | /feed/entry/content                                                           | /items/content_html                  |
| Link          | /rss/channel/item/link<br>/rdf: RDF/item/link                                                                                                                                       | /feed/entry/link[@rel=”alternate”]/@href<br>/feed/entry/link[not(@rel)]/@href | /items/url                           |
| Updated       | /rss/channel/item/dc:date<br>/rdf: RDF/rdf:item/dc:date                                                                                                                             | /feed/entry/modified<br>/feed/entry/updated                                   | /items/date_modified                 |
| Published     | /rss/channel/item/pubDate<br>/rss/channel/item/dc:date                                                                                                                              | /feed/entry/published<br>/feed/entry/issued                                   | /items/date_published                |
| Author        | /rss/channel/item/author<br>/rss/channel/item/dc:author<br>/rdf: RDF/item/dc:author<br>/rss/channel/item/dc:creator<br>/rdf: RDF/item/dc:creator<br>/rss/channel/item/itunes:author | /feed/entry/author                                                            | /items/author/name                   |
| Authors       | /rss/channel/item/author<br>/rss/channel/item/dc:author<br>/rdf: RDF/item/dc:author<br>/rss/channel/item/dc:creator<br>/rdf: RDF/item/dc:creator<br>/rss/channel/item/itunes:author | /feed/entry/authors[0]                                                        | /items/authors<br>/items/author/name |
| GUID          | /rss/channel/item/guid                                                                                                                                                              | /feed/entry/id                                                                | /items/id                            |
| Image         | /rss/channel/item/itunes:image<br>/rss/channel/item/media:image                                                                                                                     |                                                                               | /items/image<br>/items/banner_image  |
| Categories    | /rss/channel/item/category<br>/rss/channel/item/dc:subject<br>/rss/channel/item/itunes:keywords<br>/rdf: RDF/channel/item/dc:subject                                                | /feed/entry/category                                                          | /items/tags                          |
| Enclosures    | /rss/channel/item/enclosure                                                                                                                                                         | /feed/entry/link[@rel=”enclosure”]                                            | /items/attachments                   |

Added a noescape function to templates, that prevents escaping of HTML.

```
{{ .Content|noescape}}
```



## Changelog

### 10.01.2023 - [Alex](github.com/portalzine)

- Updated styles
- Updated and cleaned up the [Go](https://go.dev/) feed render script
- Moved feed rendering to a template for easy tweaking
- Moved to [Font-Awesome](https://fontawesome.com) 6 Free
- Documentation update
- Add noescape function for HTML

### 23.03.2017 - [Erik](https://github.com/edasque)

- Initial Release
