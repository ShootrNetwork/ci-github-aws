package main

import (
	"fmt"
	"log"
	"time"
)

func uploadArtifactsToS3(params Params) {

	branchCheck := BranchCheck{Params: params}

	if branchCheck.should_execute_upload_to_s3() {
		start := time.Now()
		log.Println("Upload to s3 start...")

		commit := params.Git.Commit
		aws := params.Config.AWS
		InitAWSSession(aws.Region)
		bucket := aws.ActifactBucket
		//listBuckets()

		uploadToS3(bucket, "./shootr-api/target/shootr-api.jar", fmt.Sprintf("artifacts/shootr-api.jar.%s", commit))
		uploadToS3(bucket, "./shootr-services/target/shootr-services.jar", fmt.Sprintf("artifacts/shootr-services.jar.%s", commit))
		uploadToS3(bucket, "./shootr-backoffice/target/shootr-backoffice.jar", fmt.Sprintf("artifacts/shootr-backoffice.jar.%s", commit))

		log.Printf("Upload to s3 done in %s", time.Since(start))

	} else {
		log.Println("Skipping Upload to s3")
	}
}
