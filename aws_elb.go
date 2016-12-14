package main

import (
	"errors"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/elb"
)

func elbCheckInstancesInService(elbName string) error {
	svc := elb.New(awsSession)

	response, err := svc.DescribeInstanceHealth(&elb.DescribeInstanceHealthInput{
		LoadBalancerName: aws.String(elbName),
	})

	if err == nil {
		for _, state := range response.InstanceStates {
			if *state.State != "InService" {
				log.Printf("instance %s -> %s", *state.InstanceId, *state.State)
				return errors.New("not ready yet")
			}
		}
	}

	return err
}
