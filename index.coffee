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
  font-family: Helvetica
  left: 0px
  top: 30px
  font-size: .8em

  A
    color:white
 
  h3
    font-style:"Futura"
    font-weight: 500
    margin: 0
    padding: 0 0 0 5px
  """

