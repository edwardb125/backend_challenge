package Struct

import (
"context"
"errors"
"github.com/aws/aws-sdk-go/service/dynamodb"
"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
"github.com/stretchr/testify/assert"
"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		name        string
		marshalErr  error
		dynamoERR   error
		errExpected error
		output      Output
	}{
		{name: "marshalErr", marshalErr: errors.New("error"), errExpected: errors.New("server error"), output: Output{}},
		{name: "dynamoErr", dynamoERR: errors.New("error"), errExpected: errors.New("server error"), output: Output{}},
		{name: "ok", output: Output{Message: "request done successfully"}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
				marshalMock := func(in interface{}) (map[string]*dynamodb.AttributeValue, error) {
					return map[string]*dynamodb.AttributeValue{}, test.marshalErr
				}
				dynamo := mockDynamo{
					expectedErr: test.dynamoERR,
				}
				core := Core{
					DyDB:       &dynamo,
					MarshalMap: marshalMock,
				}

				handler, err := core.Handler(context.TODO(), Input{})

				assert.Equal(t, test.output, handler)

				if err == nil {
					assert.Nil(t, test.errExpected)
				} else {
					assert.EqualError(t, err, test.errExpected.Error())
				}

				})
				}
	}

type mockDynamo struct {
	dynamodbiface.DynamoDBAPI
	expectedErr error
}

func (m *mockDynamo) PutItem(*dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	return &dynamodb.PutItemOutput{}, m.expectedErr
}