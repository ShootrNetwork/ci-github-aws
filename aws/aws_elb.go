package aws

import (
	"errors"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/elb"
)

func ElbCheckInstancesInService(elbName string) error {
	svc := elb.New(awsSession)

	response, err := svc.DescribeInstanceHealth(&elb.DescribeInstanceHealthInput{
		LoadBalancerName: aws.String(elbName),
	})

	if err == nil {
		for _, state := range response.InstanceStates {
			if *state.State != "InService" {
				log.Printf("instance %s -> %s", *state.InstanceId, *state.State)
				return errors.New("Not ready yet")
			}
		}
	}
	return err
}

func GetElb(elbName string) *elb.LoadBalancerDescription {
	svc := elb.New(awsSession)

	response, err := svc.DescribeLoadBalancers(&elb.DescribeLoadBalancersInput{
		LoadBalancerNames: []*string{aws.String(elbName)},
	})
	check(err)

	return response.LoadBalancerDescriptions[0]
}

func ElbCheckInstanceCountIsDesired(elbName string, desired int) error {
	elb := GetElb(elbName)
	currentCount := len(elb.Instances)
	if currentCount != desired {
		log.Printf("ELB Instances/desired -> (%d/%d), ids: %v", currentCount, desired, elb.Instances)
		return errors.New("Not ready yet")
	}
	return nil
}

func ElbRemoveInstances(elbName string, instanceIds []*string) {
	svc := elb.New(awsSession)

	var elbInstances []*elb.Instance
	for _, instanceId := range instanceIds {
		elbInstance := &elb.Instance{
			InstanceId: aws.String(*instanceId),
		}
		elbInstances = append(elbInstances, elbInstance)
	}

	svc.DeregisterInstancesFromLoadBalancer(&elb.DeregisterInstancesFromLoadBalancerInput{
		LoadBalancerName: aws.String(elbName),
		Instances:        elbInstances,
	})
}
