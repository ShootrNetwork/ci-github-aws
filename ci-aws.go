package main

import "log"

func main() {
	log.Println("Starting...")

	params := parseParams()
	log.Printf("Params: %+v\n", params)

	switch params.Command {

	case "test_and_build":
		testAndBuild(params)

	case "upload_to_s3":
		uploadArtifactsToS3(params)

	case "docker_build":
		dockerBuildComponents(params)

	case "docker_tag":
	case "deploy":
	default:
		panic("unrecognized command")
	}
}
