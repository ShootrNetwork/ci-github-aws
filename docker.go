package main

import (
	"fmt"
	"log"
	"time"
)

func dockerBuild(path string, dockerfile string, imageName string) {
	log.Printf("building docker %s", imageName)
	command := fmt.Sprintf("docker build -f %s -t %s %s", dockerfile, imageName, path)
	exe_cmd_wait(command, 5*time.Minute)
	log.Printf("build command: %s", command)
}

func dockerPush(imageName string) {
	log.Printf("pushing docker image %s", imageName)
	command := fmt.Sprintf("docker push %s", imageName)
	exe_cmd_wait(command, 10*time.Minute)
	log.Printf("push command: %s", command)
}
