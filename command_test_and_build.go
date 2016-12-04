package main

import (
	"log"
	"time"
)

func TestAndBuild(params Params) {
	start := time.Now()
	log.Println("Test and build start...")

	command := "docker exec java8-ci mvn --quiet --batch-mode -f /ci/pom.xml clean test install jacoco:report coveralls:report"
	exe_cmd_wait(command)

	log.Printf("Test and build done in %s", time.Since(start))
}
