package main

import (
  "fmt"
)


func main() {
  // bucket := "image-repo-manan"

 //  tags1 := []string {"art", "sports", "music"}
 //  tags2 := []string {"sports", "drama"}
 //  tags3 := []string {"edu", "bank"}
 //  tags4 := []string {"press", "movie"}
 //  img1 := CreateImage("Manan", "first-image","just testing","first-image", tags1, tags1)
 //  img2 := CreateImage("Bhavya", "second-image", "just testing", "second-image", tags2, tags2)
 //  img3 := CreateImage("Jags", "third-image", "just testing", "third-image", tags3, tags3)
 //  img4 := CreateImage("Kiran", "fourth-image", "just testing", "fourth-image",tags4, tags4)
 //  imgs := []*Image {img1, img2, img3, img4}
	// es, err := NewElasticClient("image-repo")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// res ,err := es.AutomcompleteSuggesterText("img")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(res)
	s3bucket := "image-repo-manan"
	esIndex := "image-repo"
	rp, er := NewRepo(s3bucket, esIndex)
	if er != nil {
		fmt.Println(er)
	}
	tags1 := []string {"art", "sports", "music"}
	_, err := rp.AddImage("./mkm1.jpg", "Manan", "first-image","just testing","first-image", tags1, tags1)
	if err != nil {
		fmt.Println(err)
	}

	th := []string {"iJ_2pCaZl79XjQ7jvxuJKFd0bCWJIZ-7wH1hoAmHM58"}

	fields := []string {"Text"}

	res, err2 := rp.es.MoreLikeThis(th, fields)

	if err != nil {
		fmt.Println(err2)
	}

	fmt.Println(res)

	// for i := 0; i < len(imgs); i++ {
	// 	er := es.AddDoc(imgs[i])
	// 	if er != nil {
	// 		fmt.Println(er)
	// 	}
	// }
	// // // h4 := img4.Hash

	// // err1 := es.AddDoc(img4)

	// // if err1 != nil {fmt.Println(err)}

	// th := []string {"iJ_2pCaZl79XjQ7jvxuJKFd0bCWJIZ-7wH1hoAmHM58"}

	// // tags := []string {"sports", "movie"}
	// fields := []string {"Text"}
	// res, err := es.MoreLikeThis(th, fields)

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(res)
}

