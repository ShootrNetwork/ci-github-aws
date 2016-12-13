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

func dockerPull(imageName string) {
	log.Printf("pulling docker image %s", imageName)
	command := fmt.Sprintf("docker pull %s", imageName)
	exe_cmd_wait(command, 10*time.Minute)
	log.Printf("pull command: %s", command)
}

func dockerTag(from string, to string) {
	log.Printf("taging docker image %s to %s", from, to)
	command := fmt.Sprintf("docker tag %s %s", from, to)
	exe_cmd_wait(command, 2*time.Minute)
	log.Printf("tag command: %s", command)
}
