# Image-repo (Go)

A sample implementation of an possible image-repository or storage. Focusing more on the search functionality using Elasticsearch and for storages purposes I am leveraging AWS S3.

## Installation

You need to have Elasticsearch version >= `7.0` installed and shoyuld have your aws config and creadentials in `~/.aws/config` and `~/.aws/credentials` respectively. 

```sh
-> go get -v github.com/olivere/elastic
-> go get -v github.com/aws/aws-sdk-go
-> go run *.go
```

## Example/Usage

**Creating and storing an Image**

```go
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
 ``` 
 
 **Search Functionality**
 
 The search functions currently supported are 
 ```go
 func (this *Repo) Exists(hash string) error
 func (this *Repo) SearchWithTerm (k string, v interface{})
 func (this *Repo) SearchWithTags (tags []string)
 func (this *Repo) MoreLikeThis (hashes, fields []string)
 func (this *Repo) AutomcompleteSuggester (prefix, field string)
 ```
 
 ## Author

Manan Maniyar:[ E-mail](mailto:maniyarmanan1996@gmail.com), [@Manan007224](https://www.github.com/Manan007224)
