package main

type Repo struct {
	s3	*S3Client
	es	*Elastic
	bucket, index string
} 

func NewRepo (bck, idx string) (*Repo, error) {
	svc := S3Init(bck)
	esc, err := NewElasticClient(idx)
	if err != nil {
		return nil, err
	}
	r := &Repo {
		s3: svc,
		es: esc,
		bucket: bck,
		index: idx,
	}
	return r, nil
}

// AddImage function does three things
// 1) - Create an Image struct
// 2) - Uplaod it to S3
// 3) - Upload it to Elasticsearch

func (this *Repo) AddImage(filepath, user, text, desc, suggestt string, tags, suggestc []string) (*Image, error) {
	img := CreateImage(user, text, desc, suggestt, tags, suggestc)
	img.S3key = filepath
	err := this.s3.Upload(filepath)
	if err != nil {
		return img, err
	}
	er := this.es.AddDoc(img)
	if er != nil {
		return img, er
	}
	return img, nil
}

// Search Functions

func (this *Repo) Exists(hash string) error {
	return this.es.Exists(hash)
}

func (this *Repo) SearchWithTerm (k string, v interface{}) ([]Image, error) {
	return this.es.SearchWithTerm(k, v)
}

func (this *Repo) SearchWithTags (tags []string) ([]Image, error) {
	return this.es.SearchWithTags(tags)
}

func (this *Repo) MoreLikeThis (hashes, fields []string) ([]Image, error) {
	return this.es.MoreLikeThis(hashes, fields)
}

func (this *Repo) AutomcompleteSuggester (prefix, field string) ([]string, error) {
	return this.es.AutomcompleteSuggester(prefix, field)
}