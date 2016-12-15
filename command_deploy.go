package main

import (
	"fmt"
	"log"
	"time"

	awsShootr "github.com/shootrnetwork/ci-github-aws/aws"
)

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
			deployASG(asg)
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

func deployASG(asgName string) {
	asg := awsShootr.GetASG(asgName)
	instanceIds := awsShootr.GetASGInstanceIdsFromGroup(asg)
	log.Printf("ASG Instances: %v", instanceIds)

	oldDesiredCapacity := asg.DesiredCapacity
	newDesiredCapacity := *oldDesiredCapacity * int64(2)
	log.Printf("Setting ASG desired capacity from %d to %d", oldDesiredCapacity, newDesiredCapacity)
	awsShootr.SetASGDesiredCapacity(asgName, newDesiredCapacity)

	const timeout = 10 * time.Minute
	const timeBetweenExecutions = 10 * time.Second

	log.Println("Checking for instance count in ASG to be ok")
	executeWithTimeout(timeout, timeBetweenExecutions, func() error {
		return awsShootr.AsgCheckInstanceCountIsDesired(asgName)
	})

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

	log.Printf("Setting ASG desired capacity back to %d", oldDesiredCapacity)
	awsShootr.SetASGDesiredCapacity(asgName, *oldDesiredCapacity)

	log.Println("Checking for instance count in ASG to be ok")
	executeWithTimeout(timeout, timeBetweenExecutions, func() error {
		return awsShootr.AsgCheckInstanceCountIsDesired(asgName)
	})
}
