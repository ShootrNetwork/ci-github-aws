package main

import (
	"fmt"
	"log"
	"time"
)

func testAndBuild(params Params) {
	branchCheck := BranchCheck{Params: params}

	if branchCheck.should_execute_test_and_build() {
		start := time.Now()
		log.Println("Test and build start...")

		//command := fmt.Sprintf("docker exec java8-ci mvn --quiet --batch-mode -f %s/pom.xml clean test install jacoco:report coveralls:report", params.Config.PathInDocker)
		//command := fmt.Sprintf("docker exec java8-ci mvn --quiet --batch-mode -f %s/pom.xml clean", params.Config.PathInDocker)
		command := fmt.Sprintf("docker exec java8-ci mvn -f %s/pom.xml clean test install", params.Config.PathInDocker)
		exe_cmd_wait(command, 15*time.Minute)

		log.Printf("Test and build done in %s", time.Since(start))

	} else {
		log.Println("Skipping Test and build")
	}
}
