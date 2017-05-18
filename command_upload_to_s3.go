package main

import (
	"fmt"
	"log"
	"time"

	awsShootr "github.com/shootrnetwork/ci-github-aws/aws"
)

func uploadArtifactsToS3(params Params) {

	branchCheck := BranchCheck{params.Config.CurrentConfig}

	if branchCheck.shouldExecuteUploadToS3() {
		start := time.Now()
		log.Println("Upload to s3 start...")

		commit := params.Git.Commit
		aws := params.Config.AWS
		bucket := aws.ActifactBucket

		awsShootr.InitAWSSession(aws.Region)

		for _, component := range params.Config.Components {
			localFile := fmt.Sprintf("%s/%s", component.JarPath, component.JarName)
			s3File := fmt.Sprintf("%s/%s.%s", aws.ArtifactFolder, component.JarName, commit)
			awsShootr.UploadToS3(bucket, localFile, s3File)
		}

		log.Printf("Upload to s3 done in %s", time.Since(start))

	} else {
		log.Println("Skipping Upload to s3")
	}
}
