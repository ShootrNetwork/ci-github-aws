package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

var awsSession *session.Session

func ConfigureAWS(region string) {
	awsSession = session.New()
	cred := credentials.NewChainCredentials([]credentials.Provider{
		&credentials.EnvProvider{},
	})
	awsSession.Config.Credentials = cred
	awsSession.Config.Region = aws.String(region)
	awsSession.Config.WithMaxRetries(aws.UseServiceDefaultRetries)
}
