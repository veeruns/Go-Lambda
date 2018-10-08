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
