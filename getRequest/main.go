package getRequest

import (
	"backendChallenge/getRequest/Struct"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"log"
	"os"
)

func main(){
	region := os.Getenv("AWS_REGION")
	newSession, errorMessage := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})

	if errorMessage != nil{
		log.Println(errorMessage)
		return
	}
	// create new dynamoDB with session data
	dbClient := dynamodb.New(newSession)
	// bottom line use core struct in struct directory to make new object from that
	core := Struct.Core{
		DyDB:        dbClient,
	}
	lambda.Start(core.Handler)
}
