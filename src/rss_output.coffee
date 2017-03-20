FeedParser = require("feedparser")
request = require("request")
config = require('../config.json');

# req = request("http://rss.cnn.com/rss/cnn_topstories.rss")
req = request("http://newsrss.bbc.co.uk/rss/newsonline_world_edition/front_page/rss.xml")
max_articles = 6

console.log("<h3>BBC News Top Stories</h3>")
console.log(config.feeds.length)
console.log ("<ul>")

feedparser = new FeedParser()

req.on "error", (error) ->


# handle any request errors
req.on "response", (res) ->
  stream = this
  return @emit("error", new Error("Bad status code"))  unless res.statusCode is 200
  stream.pipe feedparser
  return

feedparser.on "error", (error) ->
# always handle errors

feedparser.on "end", (error) ->
  console.log ("</ul>")



feedparser.on "readable", ->
  
  # This is where the action is!
  stream = this
  meta = @meta # **NOTE** the "meta" is always available in the context of the feedparser instance
  item = undefined
  while item = stream.read()
    console.log "<li>#{item.title} - <a href='#{item.link}''> <i class='fa fa-external-link'></i> </a></li>" if (max_articles>0)
    max_articles--
  return
