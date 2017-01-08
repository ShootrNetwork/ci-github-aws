package main

import (
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/service/autoscaling"
	awsShootr "github.com/shootrnetwork/ci-github-aws/aws"
)

const timeout = 10 * time.Minute
const timeBetweenExecutions = 10 * time.Second

func deployComponents(params Params) {

	branchCheck := BranchCheck{params.Config.CurrentConfig}

	if branchCheck.shouldDeploy() {
		start := time.Now()
		log.Println("Deploy start...")

		backofficeUrl := branchCheck.getBackofficeUrl()
		if backofficeUrl != "" {
			deployBackoffice(backofficeUrl, params.Pem)
		}

		asg := branchCheck.getASG()
		if asg != "" {
			awsShootr.InitAWSSession(params.Config.AWS.Region)
			deployASG(asg, branchCheck.getDeployType())
		}

		log.Printf("Deploy done in %s", time.Since(start))

	} else {
		log.Println("Skipping Deploy")
	}
}

func deployBackoffice(url string, pem string) {
	log.Printf("deploying backoffice: %s", url)
	command := fmt.Sprintf("ssh -oStrictHostKeyChecking=no -i %s ubuntu@%s /home/ubuntu/set_env_and_start.sh", pem, url)
	exe_cmd_wait(command, 10*time.Minute)
	log.Printf("deploying backoffice: %s", command)
}

func deployASG(asgName string, deployType string) {
	asg := awsShootr.GetASG(asgName)
	instanceIds := awsShootr.GetASGInstanceIdsFromGroup(asg)
	log.Printf("ASG Instances: %+v", &instanceIds)

	oldDesiredCapacity := int(*asg.DesiredCapacity)
	if oldDesiredCapacity < 1 {
		log.Fatalf("ASG %s has no instances!!!", asgName)
	}

	newDesiredCapacity := oldDesiredCapacity * 2
	log.Printf("Setting ASG desired capacity from %d to %d", oldDesiredCapacity, newDesiredCapacity)
	awsShootr.SetASGDesiredCapacity(asgName, int64(newDesiredCapacity))

	log.Println("Checking for instance count in ASG to be ok")
	executeWithTimeout(timeout, timeBetweenExecutions, func() error {
		return awsShootr.AsgCheckInstanceCountIsDesired(asgName)
	})

	if deployType == "elb" {
		doElbChecks(asg, newDesiredCapacity, instanceIds)
	} else if deployType == "targetGroup" {
		doTargetGroupChecks(asg, newDesiredCapacity, instanceIds)
	}

	log.Printf("Setting ASG desired capacity back to %d", oldDesiredCapacity)
	awsShootr.SetASGDesiredCapacity(asgName, int64(oldDesiredCapacity))

	log.Println("Checking for instance count in ASG to be ok")
	executeWithTimeout(timeout, timeBetweenExecutions, func() error {
		return awsShootr.AsgCheckInstanceCountIsDesired(asgName)
	})
}

func doElbChecks(asg *autoscaling.Group, newDesiredCapacity int, instanceIds []*string) {
	elbName := asg.LoadBalancerNames[0]
	log.Println("Checking for instance count in ELB to be ok")
	executeWithTimeout(timeout, timeBetweenExecutions, func() error {
		return awsShootr.ElbCheckInstanceCountIsDesired(*elbName, int(newDesiredCapacity))
	})

	log.Println("Waiting for all ELB instances to be healthy")
	executeWithTimeout(timeout, timeBetweenExecutions, func() error {
		return awsShootr.ElbCheckInstancesInService(*elbName)
	})

	log.Printf("Removing original instances from ELB to drain connections: %v", instanceIds)
	awsShootr.ElbRemoveInstances(*elbName, instanceIds)
	time.Sleep(10 * time.Second)
}

func doTargetGroupChecks(asg *autoscaling.Group, newDesiredCapacity int, instanceIds []*string) {
	targetGroupArn := asg.TargetGroupARNs[0]
	log.Println("Checking for instance count in TargetGroup to be ok")
	executeWithTimeout(timeout, timeBetweenExecutions, func() error {
		return awsShootr.TargetGroupCheckInstanceCountIsDesired(*targetGroupArn, int(newDesiredCapacity))
	})

	log.Println("Waiting for all Target Group instances to be healthy")
	executeWithTimeout(timeout, timeBetweenExecutions, func() error {
		return awsShootr.AlbCheckInstancesInService(*targetGroupArn)
	})

	log.Printf("Removing original instances from ALB to drain connections: %v", instanceIds)
	awsShootr.AlbRemoveInstances(*targetGroupArn, instanceIds)
	time.Sleep(10 * time.Second)
}
