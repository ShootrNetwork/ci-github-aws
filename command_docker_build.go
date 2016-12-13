package main

import (
	"fmt"
	"log"
	"time"
)

func dockerBuildComponents(params Params) {
	branchCheck := BranchCheck{Params: params}

	if branchCheck.should_execute_test_and_build() {
		start := time.Now()
		log.Println("Docker build start...")

		commit := params.Git.Commit

		buildAndPushDocker("api", commit)
		buildAndPushDocker("services", commit)
		buildAndPushDocker("backoffice", commit)

		log.Printf("Test and build done in %s", time.Since(start))

	} else {
		log.Println("Skipping Test and build")
	}
}

func buildAndPushDocker(component string, commit string) {
	copy_jar(component)

	image := fmt.Sprintf("fav24/shootr-%s:%s", component, commit)
	path := fmt.Sprintf("./shootr-%s", component)
	dockerfile := path + "/Dockerfile"

	dockerBuild(path, dockerfile, image)
	dockerPush(image)
}

func copy_jar(component string) {
	from := fmt.Sprintf("./shootr-%s/target/shootr-%s.jar", component, component)
	to := fmt.Sprintf("./shootr-%s/shootr-%s.jar", component, component)
	log.Printf("Copying file from %s to %s", from, to)
	copyFile(from, to)
}
