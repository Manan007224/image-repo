package main

import (
  "fmt"
)


func main() {
  // bucket := "image-repo-manan"
  // client := S3Init("image-repo-manan")
  // err := client.S3Upload("mkm.jpg")
  // if err != nil {
  // 	fmt.Println(err)
  // }
  // tags1 := []string {"art", "sports", "music"}
  // tags2 := []string {"sports", "drama"}
  // tags3 := []string {"edu", "bank"}
  // img1 := CreateImage("Manan", "first-image", "just testing", tags1)
  // img2 := CreateImage("Bhavya", "second-image", "just testing", tags2)
  // img3 := CreateImage("Jags", "third-image", "just testing", tags3)

	es, err := NewElasticClient("images")
	if err != nil {
		fmt.Println(err)
	}
	// tags4 := []string {"press", "movie"}
	// img4 := CreateImage("Kiran", "fourth-image", "just testing", tags4)
	
	// // h4 := img4.Hash

	// err1 := es.AddDoc(img4)

	// if err1 != nil {fmt.Println(err)}

	th := []string {"iJ_2pCaZl79XjQ7jvxuJKFd0bCWJIZ-7wH1hoAmHM58"}

	// tags := []string {"sports", "movie"}
	fields := []string {"Text"}
	res, err := es.MoreLikeThis(th, fields)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)
}

