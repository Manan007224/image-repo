package Models

type Image struct {
	User								string			`json:"user"`
	Categories					[]string		`json:charachteristics"`
	Text								string			`json:text`
	S3key								string			`json:s3key`
	Hash								string			`json:hash`
}

const svc := Aws.S3Client.NewS3Client()

func (img *Image) (user, text string, tags Categories) {
	
}