package main
import (
"github.com/aws/aws-sdk-go/aws"
   "github.com/aws/aws-sdk-go/aws/session"
   "github.com/aws/aws-sdk-go/service/dynamodb"
   "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)
//US east 1 is where the data for capitals is
var db = dynamodb.New(session.New(),aws.NewConfig().WithRegion("us-east-1"))
