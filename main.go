package main

import (
	"github.com/olivere/elastic"
	"fmt"
	"context"
)

type Image struct {
	User								string			`json:"user"`
	Charachteristics		[]string		`json:charachteristics"`
	Text								string			`json:text`
	S3key								string			`json:s3key`
}


func main() {
	client, err := elastic.NewClient()
	if err != nil {
		fmt.Println("error in connection")
	}

	// fmt.Println(client.c.urls)

	// Create an index
	_, err = client.CreateIndex("tweets").Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}
}