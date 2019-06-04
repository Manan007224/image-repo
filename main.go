package main

import(  
  // "github.com/aws/aws-sdk-go/aws/credentials"
  // "github.com/aws/aws-sdk-go/aws"
  // "github.com/aws/aws-sdk-go/aws/session"
  // "github.com/aws/aws-sdk-go/service/s3"    
  "fmt"
  // "reflect" 
) 

// func main() {
// 	client, err := elastic.NewClient()
// 	if err != nil {
// 		fmt.Println("error in connection")
// 	}

// 	// fmt.Println(client.c.urls)

// 	// Create an index
// 	_, err = client.CreateIndex("tweets").Do(context.Background())
// 	if err != nil {
// 		// Handle error
// 		panic(err)
// 	}
// }

// func main() {
// 	awsAccessKey := "AKIAJF2LHYHSQEFMY5YQ"
// 	awsSecret := "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"
// 	token := ""

// 	creds := credentials.NewStaticCredentials(awsAccessKey, awsSecret, token)

// 	_, err := creds.Get()

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	cfg := aws.NewConfig().WithRegion("us-west-1").WithCredentials(creds) 
//   svc := s3.New(session.New(), cfg)

//   fmt.Println(reflect.TypeOf(svc)) 
// }

func main() {
	svc, err := S3Client.New("images")
	if err != nil {
		fmt.Println(err)
	}
	svc.Upload("mkm.jpg")
}