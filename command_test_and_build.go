package main

import (
	"log"
	"time"
)

func TestAndBuild(params Params) {
	branchCheck := BranchCheck{Params: params}

	if branchCheck.should_execute_test_and_build() {
		start := time.Now()
		log.Println("Test and build start...")

		//command := "docker exec java8-ci mvn --quiet --batch-mode -f /ci/pom.xml clean test install jacoco:report coveralls:report"
		command := "docker exec java8-ci mvn --quiet --batch-mode -f /ci/pom.xml clean"
		exe_cmd_wait(command)

		log.Printf("Test and build done in %s", time.Since(start))

	} else {
		log.Println("Skipping Test and build")
	}
}
