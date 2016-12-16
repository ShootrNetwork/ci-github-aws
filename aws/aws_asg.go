package aws

import (
	"errors"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/autoscaling"
)

func GetASG(asgName string) *autoscaling.Group {
	svc := autoscaling.New(awsSession)

	log.Printf("Request to get Asg with name: %s", asgName)

	resp, err := svc.DescribeAutoScalingGroups(&autoscaling.DescribeAutoScalingGroupsInput{
		AutoScalingGroupNames: []*string{
			aws.String(asgName)},
	})
	check(err)

	if len(resp.AutoScalingGroups) < 1 {
		log.Fatalf("ASG '%s' not found!", asgName)
	}

	return resp.AutoScalingGroups[0]
}

func GetASGInstanceIds(asgName string) []*string {
	asg := GetASG(asgName)
	return GetASGInstanceIdsFromGroup(asg)
}

func GetASGInstanceIdsFromGroup(group *autoscaling.Group) []*string {
	instances := group.Instances
	ids := make([]*string, len(instances))
	for i, instance := range instances {
		ids[i] = instance.InstanceId
	}
	return ids
}

func SetASGDesiredCapacity(asgName string, newDesiredCapacity int64) {
	svc := autoscaling.New(awsSession)
	resp, err := svc.SetDesiredCapacity(&autoscaling.SetDesiredCapacityInput{
		AutoScalingGroupName: aws.String(asgName),
		DesiredCapacity:      aws.Int64(newDesiredCapacity),
	})
	check(err)
	log.Printf("asg SetDesiredCapacity response: %v", resp)
}

func AsgCheckInstanceCountIsDesired(asgName string) error {
	asg := GetASG(asgName)
	ids := GetASGInstanceIdsFromGroup(asg)
	desired := int(*asg.DesiredCapacity)

	if desired != len(ids) {
		log.Printf("ASG Instances/desired -> (%d/%d), ids: %v", len(ids), desired, ids)
		return errors.New("Not ready yet")
	}
	return nil
}
