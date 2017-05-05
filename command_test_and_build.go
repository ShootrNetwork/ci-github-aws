package main

import (
	"fmt"
	"log"
	"time"
)

func testAndBuild(params Params) {
	branchCheck := BranchCheck{params.Config.CurrentConfig}

	if branchCheck.shouldExecuteTestAndBuild() {
		start := time.Now()
		log.Println("Test and build start...")

		path := params.Config.PathInDocker
		mvnArgs := params.Config.MvnBuildArgs

		command := fmt.Sprintf("docker exec %s bash -c \"cd %s && ./mvnw --quiet --batch-mode clean test install %s -Dmaven.javadoc.skip=true\"", params.Config.DockerBuildImageName, path, mvnArgs)
		//command := fmt.Sprintf("docker exec java8-ci mvn --quiet --batch-mode -f %s/pom.xml clean", params.Config.PathInDocker)
		//command := fmt.Sprintf("docker exec java8-ci mvn -f %s/pom.xml clean test install", params.Config.PathInDocker)
		exe_cmd_wait(command, 15*time.Minute)

		log.Printf("Test and build done in %s", time.Since(start))

	} else {
		log.Println("Skipping Test and build")
	}
}
