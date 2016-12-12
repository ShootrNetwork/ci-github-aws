package main

import (
	"bytes"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func uploadToS3(source string, destination string, bucketName string, region string) {

	file, err := os.Open(source)
	if err != nil {
		log.Fatal("Failed to open file", err)
	}
	defer file.Close()

	fileInfo, _ := file.Stat()
	var size int64 = fileInfo.Size()
	buffer := make([]byte, size)
	file.Read(buffer)
	fileBytes := bytes.NewReader(buffer)

	svc := s3.New(session.New(&aws.Config{Region: aws.String(region)}))

	resp, err := svc.PutObject(&s3.PutObjectInput{
		Bucket:        aws.String(bucketName),
		Key:           aws.String(destination),
		Body:          fileBytes,
		ContentLength: &size,
	})

	if err != nil {
		log.Fatal("Failed to upload to s3", err)
	} else {
		log.Println(resp)
	}
}
