package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/autoscaling"
)

func getASGInstanceIds(asg string) []*string {
	svc := autoscaling.New(awsSession)

	params := &autoscaling.DescribeAutoScalingGroupsInput{
		AutoScalingGroupNames: []*string{
			aws.String(asg),
		},
	}
	resp, err := svc.DescribeAutoScalingGroups(params)
	check(err)

	return getInstanceIdsFromASG(resp.AutoScalingGroups[0])
}

func getInstanceIdsFromASG(group *autoscaling.Group) []*string {
	instances := group.Instances
	ids := make([]*string, len(instances))
	for i, instance := range instances {
		ids[i] = instance.InstanceId
	}
	return ids
}
