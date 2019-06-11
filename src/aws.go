package main

import (
  "fmt"
  "os"
  "path/filepath"
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/s3/s3manager"
)


type S3Client struct {
    sess      *session.Session
    svc       *s3manager.Uploader
    bucket    string
}

func S3Init (bck string) *S3Client {
  conf := aws.Config{Region: aws.String("us-east-1")}
  sess := session.New(&conf)
  svc := s3manager.NewUploader(sess)
  client := &S3Client{
    sess: sess,
    svc: svc,
    bucket: bck,
  }
  return client
}

func (this *S3Client) Upload (filename string) (error) {
  file, err := os.Open(filename)
  if err != nil {
    fmt.Println("Failed to open file", filename, err)
    return err
  }
  defer file.Close()
  res, err := this.svc.Upload(&s3manager.UploadInput{
    Bucket: aws.String(this.bucket),
    Key:    aws.String(filepath.Base(filename)),
    Body:   file,
  })
  if err != nil {
    return err
  }
  fmt.Println(res)
  return nil
} 