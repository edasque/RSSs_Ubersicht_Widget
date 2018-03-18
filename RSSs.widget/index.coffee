format = '%d %A %R'

command: "RSSs.widget/src/outPutFeed"

# the refresh frequency in milliseconds
refreshFrequency: "10m"

render: (output) -> """
  <link rel='stylesheet' href='https://maxcdn.bootstrapcdn.com/font-awesome/4.5.0/css/font-awesome.min.css'>
  #{output}
"""

style: """
  color: #FFFFFF
  font-family: -apple-system, BlinkMacSystemFont, sans-serif;
  background: rgba(#000, 0.6)

  left: 0px
  top: 10px
  font-size: .9em

  ul
    margin: 8px 0 20px 0

  A
    color:white
 
  h3
    font-style: Futura
    font-weight: 500
    margin: 0
    padding: 0 0 0 5px
  """

