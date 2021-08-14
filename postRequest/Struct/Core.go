package Struct

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"log"
	"os"
)

// Core struct help to store data easier as can by using some simple command
type Core struct {
	DyDB       dynamodbiface.DynamoDBAPI
	MarshalMap func(in interface{}) (map[string]*dynamodb.AttributeValue, error)
}

func (temp *Core) Handler(text context.Context, inf Input) (Output, error){
	device, errorMessage := temp.MarshalMap(inf)
	if errorMessage != nil {
		log.Println(errorMessage)
		return Output{}, errors.New("server error")
	}
	data := &dynamodb.PutItemInput{
		Item : device,
		TableName: aws.String(os.Getenv("TABLE_NAME")),
	}
	_ ,errorMessage = temp.DyDB.PutItem(data)
	if errorMessage != nil{
		log.Println(errorMessage)
		return Output{},errors.New("server error")
	}
	return Output{Message: "request done successfully"},nil
}

