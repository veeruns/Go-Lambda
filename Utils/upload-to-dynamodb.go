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
    Index int `json:"Index"`
    Country string `json:"Country"`
    City string `json:"City"`
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
    var index int
    for _, item := range capitalinfos {
      index=index+1
     av, err := dynamodbattribute.MarshalMap(item)
     item.Index=index
     if err != nil {
         fmt.Println("Got error marshalling map:")
         fmt.Println(err.Error())
         os.Exit(1)
     }

     // Create item in table Movies
     input := &dynamodb.PutItemInput{
         Item: av,
         TableName: aws.String("CapitalsIndex"),
     }

     _, err = svc.PutItem(input)

     if err != nil {
         fmt.Println("Got error calling PutItem:")
         fmt.Println(err.Error())
         os.Exit(1)
     }

     fmt.Println("Successfully added '",item.Country,"' (",item.City,") to With index ", index ,"  Capitals table")
 }
}
