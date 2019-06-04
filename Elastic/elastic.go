package Elastic

import (
	"github.com/olivere/elastic"
	"fmt"
	"context"
	"../Image"
)

type ELastic struct {
	c 		*elastic.Client
	index	string
}

func (e *Elastic) NewClient (*ELastic, error) {
	c, err := ELastic.NewConnection()
	if err != nil {
		return &Elastic{c, index}, err
	}
	return &Elastic{c, index}, nil
}

func (e *Elastic) AddImage(img *Models.Image) {

}

func (e *Elastic) Exists(hash string) bool {

}

func (e *Elastic) Search(term string, categories []uint8) ([]string, error) {

}

func (e *Elastic) SearchRange(term string, from, to int, categories []uint8) ([]string, error) {

}

func (e *Elastic) CategorySearch(tag uint8, from, to int) ([]string, error) {

}

func (e *Elastic) MoreLikeThis(hash string) ([]string, err) {
	
}