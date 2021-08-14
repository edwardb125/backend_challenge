package Struct

import (
	"errors"
	"testing"
)

func TestCore_Handler(t *testing.T) {
	tests := []struct {
		name		string
		input 		Input
		output		Output
		getItemErr 	error
		//notfoundErr error
		errExpected error
	} {
		{name:"getItemErr", getItemErr:errors.New("Internal Server Error"), errExpected: errors.New("Internal Server Error"),output: Output{}},
		{name:"unmarshalErr",  errExpected: errors.New("Internal Server Error"),output: Output{}},
	}
	for _,test :=range tests {
		t.Run(test.name, func(t *testing.T){
			//getItemMock :=func (DynamoDBAPI) GetItem(*dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error)
		})
	}
}
