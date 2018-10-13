package httpsServer
import (
  "net/http"
  "fmt"
  "bytes"
)

func PowerOff(hostname string) bool {
  var buff bytes.Buffer
  var url bytes.Buffer
  url.WriteString("http://")
  url.WriteString(hostname)
  url.WriteString("/keypress/poweron")
  resp,err := http.Post(url.String(),"",&buff)
  if(err != nil){
    fmt.Printf("That did not work as intended %s\n",err.Error())
    return false
  } else{
    fmt.Printf("The return string is %s\n",resp.Body)
    return true
  }
}


func PowerOn(hostname string) bool {
  var buff bytes.Buffer
  var url bytes.Buffer
  url.WriteString("http://")
  url.WriteString(hostname)
  url.WriteString("/keypress/poweron")
  resp,err := http.Post(url.String(),"",&buff)
  if(err != nil){
    fmt.Printf("That did not work as intended %s\n",err.Error())
    return false
  } else{
    fmt.Printf("The return string is %s\n",resp.Body)
    return true
  }
}
