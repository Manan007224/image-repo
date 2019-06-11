package main

import (
	"github.com/olivere/elastic"
	"fmt"
	"context"
	"errors"
	"encoding/json"
)

const backgroundThreads = 5

type Elastic struct {
	client 	*elastic.Client
	b  		*elastic.BulkProcessorService
	index	string
}

func NewElasticClient (idx string) (*Elastic, error) {
	conn, err := elastic.NewClient()
	if err != nil {
		return nil, err
	}
	es := &Elastic {
		client : conn,
		index : idx,
		b: nil,
	}		
	errc := es.checkExists(idx)
	if errc != nil {return es, errc}
	es.b = es.client.BulkProcessor().Name("BackgroundWorker-1")
	return es, nil
}

func (this *Elastic) checkExists (idx string) error {
	exists, err := this.client.IndexExists(idx).Do(context.Background())
	if err != nil {
		return err
	}
	if !exists { // index already there
		return nil
	}
	_, errc := this.client.CreateIndex(idx).Do(context.Background())
	if err != nil {
		return errc
	}
	return nil
}

func (this *Elastic) AddDoc (img *Image) error {
	_, err := this.client.Index().
						Index(this.index).
						Type("_doc").
						Id(img.Hash).
						BodyJson(img).
						Refresh("wait_for").
						Do(context.Background())

	if err != nil {return err}
	return nil					
}

func (this *Elastic) BulkAddDoc(imgs []*Image) error {
	p, err := this.b.Workers(backgroundThreads).Do(context.Background())
	if err != nil {
		return err
	}
	for i := 0; i < len(imgs); i++ {
		t := imgs[i]
		r := elastic.NewBulkIndexRequest().Index(this.index).Type("_doc").Id(t.Hash).Doc(t)
		p.Add(r)
	}
	err = p.Flush()
	if err != nil {return err}
	return nil
}

func (this *Elastic) Exists (hash string) error {
	res, err := this.client.Exists().Index(this.index).Id(hash).Do(context.Background())
	if err != nil {
		return err
	}
	if res {
		return nil
	}
	return errors.New("No such document")
}

func (this *Elastic) SearchWithTerm (k string, v interface{}) ([]Image, error) {
	termQuery := elastic.NewBoolQuery().Filter(elastic.NewTermQuery(k, v))
	src, err := this.client.Search().
							Index(this.index).
							Query(termQuery).
							From(0).Size(10).
							Pretty(true).	
							Do(context.Background())

	if err != nil {
		return nil, err
	}
	result, errc := decodeSearchResultHits(src.Hits.Hits)
	return result, errc
}

func (this *Elastic) SearchWithTags (tags []string) ([]Image, error) {
	itags := convertToIf(tags)
	query := elastic.NewBoolQuery().Filter(elastic.NewTermsQuery("Tags", itags...))
	src, err := this.client.Search().
							Index(this.index).
							Query(query).
							From(0).Size(10).
							Pretty(true).
							Do(context.Background())
	if err != nil {
		return nil, err
	}
	result, err := decodeSearchResultHits(src.Hits.Hits)
	return result, err
}

// This query finds the set of arguments that are like to the given set of documents
// An example would look like this
// GET _search
// {
//   "query": {
//     "more_like_this": {
//       "fields": [
//         "Tags"
//       ],
//       "like": [
//         {
//           "_index" : "images",
//           "_id" : "iJ_2pCaZl79XjQ7jvxuJKFd0bCWJIZ-7wH1hoAmHM58"
//         }
//       ]
//     }
//   }
// }

func (this *Elastic) MoreLikeThis (hashes, fields []string) ([]Image, error) {
	query := elastic.NewMoreLikeThisQuery()
	query = query.Field(fields...)
	docs := make([]*(elastic.MoreLikeThisQueryItem), len(hashes))
	for i := 0; i < len(hashes); i++ {
		docs[i] = elastic.NewMoreLikeThisQueryItem().Index(this.index).Id(hashes[i])
	}
	query = query.LikeItems(docs...)
	src, err := this.client.Search().
							Index(this.index).
							Query(query).
							From(0).Size(10).
							Pretty(true).
							Do(context.Background())
	if err != nil {
		return nil, err
	}
	result, err := decodeSearchResultHits(src.Hits.Hits)
	return result, nil								
}

func convertToIf(tags []string) []interface{} {
	itags := make([]interface{}, len(tags))
	for i := 0; i < len(tags); i++ {
		itags[i] = tags[i]
	}
	return itags
}

func decodeSearchResultHits (hits []*elastic.SearchHit) ([]Image, error) {
	result := make([]Image, len(hits))
	if len(hits) > 0 {
		j := 0
		for _, hit := range hits {
			var t Image
			err := json.Unmarshal([]byte(hit.Source), &t)
			if err != nil {
				fmt.Println(err)
			}
			result[j] = t
			j++
		}
		return result, nil
	} else {
		return nil, errors.New("No results found")
	}
}

// Automcomplete and completion stuff

// Our Index "image-repo" has two different completion suggestors
// 1) - SuggestT - suggestor for autocompleting text
// 2) - SuggestC - suggestor for autocompleting tags
// An example of an autocomplete query looks like
//
// POST image-repo/_search
// {
//   "suggest": {
//       "text-suggest" : {
//           "prefix" : "nir",
//           "completion" : {
//               "field" : "SuggestT", 
//               "size" : 5 
//           }
//       }
//   }
// }

func (this *Elastic) AutomcompleteSuggester (prefix, field string) ([]string, error) {
	var F string
	if field == "Text" {
		F = "SuggestT"
	} else if field == "Tags" {
		F = "SuggestC"
	} else {
		return nil, errors.New("not a valid option")
	}
	textSuggester := elastic.NewCompletionSuggester("text-suggest").Fuzziness(1).Text(prefix).Field(F)
	searchSource := elastic.NewSearchSource().Suggester(textSuggester).FetchSource(false).TrackScores(true)

	src, err := this.client.Search().
							Index(this.index).
							Type("_doc").
							SearchSource(searchSource).
							Do(context.Background())
	if err != nil {
		return nil, err
	}																			
	textSuggest := src.Suggest["text-suggest"]
	var result []string
	for _, options := range textSuggest {
		for _, option := range options.Options {
			result = append(result, option.Text)
		}
	}
	return result, nil									
}