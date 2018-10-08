package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

//CapitalInfo structure
type CapitalInfo struct {
	Country string `json:"Country"`
	City    string `json:"City"`
}

type CapitalIndexInfo struct {
	Index   int    `json:"Index"`
	Country string `json:"Country"`
	City    string `json:"City"`
}

//US east 1 is where the data for capitals is
var db = dynamodb.New(session.New(), aws.NewConfig().WithRegion("us-east-1"))

func getItemIdx(index int) (*CapitalIndexInfo, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String("CapitalsIndex"),
		Key: map[string]*dynamodb.AttributeValue{
			"Index": {
				N: index,
			},
		},
	}
	result, err := db.GetItem(input)
	if err != nil {
		return nil, err
	}
	if result.Item == nil {
		return nil, nil
	}

	// The result.Item object returned has the underlying type
	// map[string]*AttributeValue. We can use the UnmarshalMap helper
	// to parse this straight into the fields of a struct. Note:
	// UnmarshalListOfMaps also exists if you are working with multiple
	// items.
	cty := new(CapitalIndexInfo)
	err = dynamodbattribute.UnmarshalMap(result.Item, cty)
	if err != nil {
		return nil, err
	}

	return cty, nil

}

func getItem(country string) (*CapitalInfo, error) {
	// Prepare the input for the query.
	input := &dynamodb.GetItemInput{
		TableName: aws.String("Capitals"),
		Key: map[string]*dynamodb.AttributeValue{
			"Country": {
				S: aws.String(country),
			},
		},
	}

	// Retrieve the item from DynamoDB. If no matching item is found
	// return nil.
	result, err := db.GetItem(input)
	if err != nil {
		return nil, err
	}
	if result.Item == nil {
		return nil, nil
	}

	// The result.Item object returned has the underlying type
	// map[string]*AttributeValue. We can use the UnmarshalMap helper
	// to parse this straight into the fields of a struct. Note:
	// UnmarshalListOfMaps also exists if you are working with multiple
	// items.
	cty := new(CapitalInfo)
	err = dynamodbattribute.UnmarshalMap(result.Item, cty)
	if err != nil {
		return nil, err
	}

	return cty, nil
}
