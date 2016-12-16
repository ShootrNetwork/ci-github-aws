package aws

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

var awsSession *session.Session

func InitAWSSession(region string) {
	awsSession = session.New()

	awsSession.Config.Credentials = credentials.NewChainCredentials([]credentials.Provider{
		&credentials.EnvProvider{},
	})

	awsSession.Config.Region = aws.String(region)
	awsSession.Config.WithMaxRetries(aws.UseServiceDefaultRetries)

	log.Println("AWS session started")
}

func check(err error) {
	if err != nil {
		panic(err.Error())
	}
}
