package main

import (
	"fmt"
	"log"
)

func main() {
	log.Println("Starting...")

	params := parseParams()
	log.Printf("Params: %+v\n", params)

	switch params.Command {

	case "test_and_build":
		//testAndBuild(params)

	case "upload_to_s3":
		uploadArtifactsToS3(params.Git.Commit, params.Config.AWS)

	case "docker_build":
	case "docker_tag":
	case "deploy":
	default:
		panic("unrecognized command")
	}
}

func uploadArtifactsToS3(commit string, aws AWS) {
	InitAWSSession(aws.Region)
	bucket := aws.ActifactBucket
	//listBuckets()

	uploadToS3(bucket, "./shootr-api/target/shootr-api.jar", fmt.Sprintf("artifacts/shootr-api.jar.%s", commit))
	uploadToS3(bucket, "./shootr-services/target/shootr-services.jar", fmt.Sprintf("artifacts/shootr-services.jar.%s", commit))
	uploadToS3(bucket, "./shootr-backoffice/target/shootr-backoffice.jar", fmt.Sprintf("artifacts/shootr-backoffice.jar.%s", commit))
}
