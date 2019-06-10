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
	User								string			`json:"user"`
	Tags								[]string		`json:charachteristics"`
	Text								string			`json:text`
	S3key								string			`json:s3key`
	Hash								string			`json:hash`
	CreatedAt						time.Time		`json:createdat`
	Description					string 			`json:description`
}

func CreateImage (user, text, desc string, tags []string) *Image {
	img := &Image {
		User: user,
		Tags: tags,
		Text: text,
		S3key: "",
		Hash: "",
		CreatedAt: time.Now(),
		Description: desc,
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