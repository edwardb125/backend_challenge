package Struct

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCore_Handler(t *testing.T) {
	tests := []struct {
		name        string
		input       Input
		output      Output
		getItemErr  error
		item        map[string]*dynamodb.AttributeValue
		errExpected error
	}{
		{name: "getItemErr", getItemErr: errors.New("Internal Server Error"), errExpected: errors.New("Internal Server Error"), output: Output{}},
		{name: "not found", item: nil, errExpected: errors.New("not found"), output: Output{}},
		{name: "ok", item: map[string]*dynamodb.AttributeValue{"id": &dynamodb.AttributeValue{S: aws.String("25")}},output: Output{Id: "25"}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			dynamoDB := &mockDynamo{
				expectedErr: test.getItemErr,
				item:        test.item,
			}
			core := Core{
				DyDB: dynamoDB,
			}

			output, err := core.Handler(context.TODO(), test.input)

			if err == nil {
				assert.Nil(t, test.errExpected)
			} else {
				assert.EqualError(t, err, test.errExpected.Error())
			}

			assert.Equal(t, test.output, output)

		})
	}
}

type mockDynamo struct {
	dynamodbiface.DynamoDBAPI
	expectedErr error
	item        map[string]*dynamodb.AttributeValue
}

func (m *mockDynamo) GetItem(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	return &dynamodb.GetItemOutput{
		Item: m.item,
	}, m.expectedErr
}