package main

import (
	"errors"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/autoscaling"
)

func getASG(asgName string) *autoscaling.Group {
	svc := autoscaling.New(awsSession)

	resp, err := svc.DescribeAutoScalingGroups(&autoscaling.DescribeAutoScalingGroupsInput{
		AutoScalingGroupNames: []*string{
			aws.String(asgName)},
	})
	check(err)

	return resp.AutoScalingGroups[0]
}

func getASGInstanceIds(asgName string) []*string {
	asg := getASG(asgName)
	return getASGInstanceIdsFromGroup(asg)
}

func getASGInstanceIdsFromGroup(group *autoscaling.Group) []*string {
	instances := group.Instances
	ids := make([]*string, len(instances))
	for i, instance := range instances {
		ids[i] = instance.InstanceId
	}
	return ids
}

func setASGDesiredCapacity(asgName string, newDesiredCapacity int64) {
	svc := autoscaling.New(awsSession)
	_, err := svc.SetDesiredCapacity(&autoscaling.SetDesiredCapacityInput{
		AutoScalingGroupName: aws.String(asgName),
		DesiredCapacity:      aws.Int64(newDesiredCapacity),
	})
	check(err)
}

func asgCheckInstanceCountIsDesired(asgName string) error {
	asg := getASG(asgName)
	ids := getASGInstanceIdsFromGroup(asg)

	if *asg.DesiredCapacity != int64(len(ids)) {
		log.Printf("ASG Instances/desired -> (%d/%d), ids: %v", len(ids), asg.DesiredCapacity, ids)
		return errors.New("Not ready yet")
	}
	return nil
}
