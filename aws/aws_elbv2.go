package aws

import (
	"errors"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/elbv2"
)

func TargetGroupCheckInstanceCountIsDesired(targetGroupArn string, desired int) error {
	svc := elbv2.New(awsSession)

	resp, err := svc.DescribeTargetHealth(&elbv2.DescribeTargetHealthInput{
		TargetGroupArn: aws.String(targetGroupArn),
	})

	if err == nil {
		currentCount := len(resp.TargetHealthDescriptions)
		if currentCount != desired {
			log.Printf("ALB TargetGroup Instances/desired -> (%d/%d), ids: %v", currentCount, desired, resp.TargetHealthDescriptions)
			return errors.New("Not ready yet")
		}
	}
	return err
}

func AlbCheckInstancesInService(targetGroupArn string) error {
	svc := elbv2.New(awsSession)

	resp, err := svc.DescribeTargetHealth(&elbv2.DescribeTargetHealthInput{
		TargetGroupArn: aws.String(targetGroupArn),
	})

	if err == nil {
		for _, targetHealthDescription := range resp.TargetHealthDescriptions {
			if *targetHealthDescription.TargetHealth.State != "healthy" {
				log.Printf("instance %s -> %s", *targetHealthDescription.Target.Id, *targetHealthDescription.TargetHealth.State)
				return errors.New("Not ready yet")
			}
		}
	}
	return err
}

func AlbRemoveInstances(targetGroupArn string, instanceIds []*string) {
	svc := elbv2.New(awsSession)

	var albTargets []*elbv2.TargetDescription
	for _, instanceId := range instanceIds {
		albTarget := &elbv2.TargetDescription{
			Id: aws.String(*instanceId),
		}
		albTargets = append(albTargets, albTarget)
	}

	svc.DeregisterTargets(&elbv2.DeregisterTargetsInput{
		TargetGroupArn: aws.String(targetGroupArn),
		Targets:        albTargets,
	})
}

func DescribeTargetGroup(groupName string) *elbv2.TargetGroup {
	svc := elbv2.New(awsSession)

	params := elbv2.DescribeTargetGroupsInput{
		Names: []*string{aws.String(groupName)},
	}

	result, err := svc.DescribeTargetGroups(&params)

	if err != nil {
		return nil
	}

	return result.TargetGroups[0]
}
