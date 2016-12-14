package main

import (
	"fmt"
	"log"
	"time"
)

func deployComponents(params Params) {

	branchCheck := BranchCheck{Params: params}

	if branchCheck.shouldDeploy() {
		start := time.Now()
		log.Println("Deploy start...")

		backofficeUrl := branchCheck.getBackofficeUrl()
		if backofficeUrl != "" {
			deployBackoffice(backofficeUrl, params.Pem)
		}

		asg := branchCheck.getASG()
		if asg != "" {
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
	asg := getASG(asgName)
	instanceIds := getASGInstanceIdsFromGroup(asg)
	log.Printf("ASG Instances: %v", instanceIds)

	oldDesiredCapacity := asg.DesiredCapacity
	newDesiredCapacity := *oldDesiredCapacity * int64(2)
	log.Printf("Setting ASG desired capacity from %d to %d", oldDesiredCapacity, newDesiredCapacity)
	setASGDesiredCapacity(asgName, newDesiredCapacity)

	checkInstanceCount(asgName)
}

func checkInstanceCount(asgName string) {
	doWithTimeout(10*time.Minute, 10*time.Second, func() error {
		return elbCheckInstancesInService(asgName)
	})
}
