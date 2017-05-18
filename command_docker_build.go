package main

import (
	"fmt"
	"log"
	"time"
)

func dockerBuildComponents(params Params) {
	branchCheck := BranchCheck{params.Config.CurrentConfig}

	if branchCheck.shouldExecuteDockerBuild() {
		start := time.Now()
		log.Println("Docker build start...")

		commit := params.Git.Commit

		for _, component := range params.Config.Components {
			buildAndPushDocker(component, commit)
		}

		log.Printf("Docker build done in %s", time.Since(start))

	} else {
		log.Println("Skipping Docker build")
	}
}

func buildAndPushDocker(component Component, commit string) {
	copyJar(component)

	image := fmt.Sprintf("%s:%s", component.DockerImage, commit)
	dockerfile := component.DockerFilePath + "/Dockerfile"

	dockerBuild(component.DockerFilePath, dockerfile, image)
	dockerPush(image)
}

func copyJar(component Component) {
	from := fmt.Sprintf("%s/%s", component.JarPath, component.JarName)
	to := fmt.Sprintf("%s/%s", component.DockerFilePath, component.JarName)
	log.Printf("Copying file from %s to %s", from, to)
	copyFile(from, to)
}
