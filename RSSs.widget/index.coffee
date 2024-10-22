format = '%d %A %R'

command: "RSSs.widget/src/outPutFeed"
 
# the refresh frequency in milliseconds
refreshFrequency: "10m"

render: (output) -> """
  <link rel='stylesheet' href='https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.2.0/css/all.min.css'>
  #{output}  
"""  
 
style: """
  font-family: Montserrat, sans-serif
  font-weight: 200
  font-size: .9em
  color: #FFFFFF 
  padding:15px 
  border:1px solid rgba(255,255,255,0.05)
  border-radius: 5px 
  width: 600px
  right: calc(50% - 300px)
  bottom: calc(80px)
  
  ul
    margin: 4px 0 8px 0
    padding:0 
    list-style: none
    margin-bottom:10px 
    max-height: 230px
    overflow-y: auto

    &::-webkit-scrollbar
     width: 5px!important    

    &::-webkit-scrollbar-track 
     background: transparent!important

    &::-webkit-scrollbar-thumb 
     background: rgba(255,255,255,0.5)!important

    &::-webkit-scrollbar-thumb:hover 
     background: #555!important

  li 
    height:28px  
    line-height: 28px
    background: rgba(255,255,255,0.02)
    margin-bottom:2px
    text-overflow: ellipsis
    white-space: nowrap;
    overflow: hidden;
    padding-left:10px
    
    span
     display: inline-block!important
     overflow: hidden
     width: 540px
   
  a
    color: rgba(255,255,255,0.6)
    float:right
    margin-right:10px
    :hover
     color: rgba(255,255,255,1)
  
  h3
    font-family: Montserrat, sans-serif
    font-size:20px 
    text-transform: uppercase 
    text-align: right
    font-weight: 200
    color: rgba(255,255,255,0.5)
    margin: 0
    padding: 0 0 0 0px
  """

