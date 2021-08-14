package Struct

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"log"
	"os"
)

// Core struct help to store data easier as can by using some simple command
type Core struct {
	DyDB       dynamodbiface.DynamoDBAPI
}

func (temp *Core) Handler(text context.Context, inf Input) (Output, error){
	getItemInput := &dynamodb.GetItemInput{
		TableName: aws.String(os.Getenv("TABLE_NAME")),
		Key: map[string]*dynamodb.AttributeValue{
			"id": &dynamodb.AttributeValue{
				S: aws.String(inf.ID),
			},
		},
	}
	item, err := temp.DyDB.GetItem(getItemInput)
	if err != nil {
		log.Println(err)
		return Output{}, errors.New("Internal Server Error") /////////
	}
	if item.Item == nil {
		log.Print(err)
		return Output{}, errors.New("not found") ////////
	}
	var device Output
	err = dynamodbattribute.UnmarshalMap(item.Item, &device) // set the information in device
	if err != nil {
		log.Println(err)
		return Output{}, errors.New("   ") //\/\/\/\/\/\/\
	}
	return device,nil /////////
}


