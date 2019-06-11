package main

import (
	"fmt"
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"io"
	"time"
)

type Image struct {
	User			string			`json:"User"`
	Tags			[]string		`json:"Tags"`
	Text			string			`json:"Text"`
	S3key			string			`json:"S3key"`
	Hash			string			`json:"Hash"`
	CreatedAt		time.Time		`json:"Createdat"`
	Description		string 			`json:"Description"`
	SuggestT		string 			`json:"SuggestT"`
	SuggestC		[]string 		`json:"SuggestC"`
}

// The mapping of the Elasticsearch client looks like
// PUT imagerepo
// {
//   "mappings": {
//     "properties": {
//       "Text" : {"type" : "text"},
//       "suggestT" : {
//         "type" : "completion"
//       },
//       "suggestC" : {
//         "type" : "completion"
//       },
//       "S3key" : {"type" : "text"},
//       "Hash" : {"type" : "text"},
//       "Description" : {"type" : "text"},
//       "user" : {"type" : "text"},
//       "Tags" : {"type": "text"},
//       "CreateAt" : {"type" : "date"}
//     }
//   }
// }

func CreateImage (user, text, desc, suggestt string, tags, suggestc []string,) *Image {
	img := &Image {
		User: user,
		Tags: tags,
		Text: text,
		S3key: "",
		Hash: "",
		CreatedAt: time.Now(),
		Description: desc,
		SuggestT: suggestt,
		SuggestC: suggestc,
	}
	img.CalculateHash()
	return img
}

func (img *Image) CalculateHash() {
	img.Hash = hashImage(img)
}

func hashImage(img *Image) string {
	h := sha256.New()
	io.WriteString(h, img.Text)
	io.WriteString(h, img.Description)
	binary.Write(h, binary.LittleEndian, img.CreatedAt.Unix())
	for _, tag := range img.Tags {
		io.WriteString(h, tag)
	}
	return base64.URLEncoding.EncodeToString(h.Sum(nil))
}

func (img *Image) VerifyHash() error {
	h := hashImage(img)
	if h != img.Hash {
		return errors.New(fmt.Sprintf("mutated hash %s vs %s", h, img.Hash))
	}
	return nil
}