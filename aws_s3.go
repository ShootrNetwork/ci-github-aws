package main

import (
	"bytes"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func listBuckets() {
	s3Client := s3.New(awsSession)

	var params *s3.ListBucketsInput
	result, err := s3Client.ListBuckets(params)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println(result)
}

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

	s3Client := s3.New(awsSession)
	resp, err := s3Client.PutObject(&s3.PutObjectInput{
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
