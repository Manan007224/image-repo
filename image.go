package Models

type Image struct {
	User								string			`json:"user"`
	Tags								[]string		`json:charachteristics"`
	Text								string			`json:text`
	S3key								string			`json:s3key`
	Hash								string			`json:hash`
	CreatedAt						time.Time		`json:createdat`
	Description					string 			`json:description`
}

func (img *Image) CreateImage(user, text, desc string, tags []string, created_at time.Time) {
	img := &{user, tags, text, "", "", created_at, desc}
	img.CalculateHash()
	return img
}

func (img *Image) CalculateHash() {
	img.Hash = hashImage(img)
}

func (img *Image) UploadToS3(path string) error {
	svc := Aws.S3Client.NewS3Client()	
	k, err := svc.Upload(path)
	if err != nil {
		return k, err
	}
	img.S3key = k
	return nil
}

func hashImage(img *Image) {
	h := sha256.New()
	io.WriteString(h, (string)(img.InfoHash))
	io.WriteString(h, img.Name)
	io.WriteString(h, img.Description)
	binary.Write(h, binary.LittleEndian, img.CreatedAt.Unix())
	for _, tag := range t.Tags {
		io.WriteString(h, tag)
	}
	return base64.URLEncoding.EncodeToString(h.Sum(nil))
}

func (img *Image) verifyHash() error {
	h := hashImage(img)
	if h != img.Hash {
		return error.new(fmt.Sprintf("mutated hash %s vs %s", h, img.Hash))
	}
}