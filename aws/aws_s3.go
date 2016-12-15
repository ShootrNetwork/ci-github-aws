package aws

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func ListBuckets() {
	s3Client := s3.New(awsSession)

	var params *s3.ListBucketsInput
	result, err := s3Client.ListBuckets(params)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println(result)
}

func UploadToS3(bucket string, source string, destination string) {

	file, err := os.Open(source)
	if err != nil {
		log.Fatal("Failed to open file", err)
	}
	defer file.Close()

	byteArray, err := ioutil.ReadAll(file)

	s3Client := s3.New(awsSession)

	resp, err := s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(destination),
		ACL:    aws.String("public-read"),
		Body:   bytes.NewReader(byteArray),
	})

	log.Printf("Uploading file to %s%s", bucket, destination)

	if err != nil {
		log.Fatal("Failed to upload to s3: ", err)
	} else {
		log.Println(resp)
	}
}
