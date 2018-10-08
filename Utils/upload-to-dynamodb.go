package main

import (
"encoding/json"
   "fmt"
   "io/ioutil"
   "os"

   "github.com/aws/aws-sdk-go/aws"
   "github.com/aws/aws-sdk-go/aws/session"
   "github.com/aws/aws-sdk-go/service/dynamodb"
   "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)


type CapitalInfo struct {
    Country string `json:"Country"`
    Capital string `json:"Capital"`
}


func Readfile() []CapitalInfo {
  raw,err := ioutil.ReadFile("../Data/country-to-capital.json")
  if(err != nil) {
    fmt.Println(err.Error())
    os.Exit(1)
  }
  var capitalinfos []CapitalInfo
  json.Unmarshal(raw, &capitalinfos)
  return capitalinfos

}

func main() {
  sess, err := session.NewSession(&aws.Config{
        Region: aws.String("us-east-1")},
    )
    if err != nil {
        fmt.Println("Error creating session:")
        fmt.Println(err.Error())
        os.Exit(1)
    }
    svc := dynamodb.New(sess)
    capitalinfos := Readfile()
    
}
