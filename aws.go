package main

import (
  "bytes"
  "log"
  "net/http"
  "os"

  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/s3"
)

type S3Client struct {
	s      *session.Session
  bucket string
}

func NewS3Client(string bucketName) (*S3Client, err) {
	awsAccessKey := "AKIAJF2LHYHSQEFMY5YQ"
	awsSecret := "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"
	token := ""

	creds := credentials.NewStaticCredentials(awsAccessKey, awsSecret, token)
	_, err := creds.Get()
	if err != nil {
		return nil, err
	}

	s, err := session.NewSession(&aws.Config{Region: aws.String(S3_REGION)})
	return &AwsClient{s, bucketName}, nil
}

func (this *S3Client) Upload(path string) string, error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	fileInfo, _ := file.Stat()
  var size int64 = fileInfo.Size()
  buffer := make([]byte, size)
  file.Read(buffer)

  _, err = s3.New(this.session).PutObject(&s3.PutObjectInput {
      Bucket:               aws.String(this.bucket),
      Key:                  aws.String(fileDir),
      ACL:                  aws.String("public"),
      Body:                 bytes.NewReader(buffer),
      ContentLength:        aws.Int64(size),
      ContentType:          aws.String(http.DetectContentType(buffer)),
      ContentDisposition:   aws.String("attachment"),
      ServerSideEncryption: aws.String("AES256"),
  })
  return path, nil
} 