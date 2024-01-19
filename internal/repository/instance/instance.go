package instance

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/credentials"
)

func GetConnection() *dynamodb.DynamoDB {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
    SharedConfigState: session.SharedConfigEnable,
    Config: aws.Config{
        Credentials: credentials.NewStaticCredentials("dummy", "dummy", "dummy"),
        Endpoint:    aws.String("http://localstack:4566"),//tpdo: from config
        Region:      aws.String("us-west-2"),
    },
	}))
	
return dynamodb.New(sess)

}
