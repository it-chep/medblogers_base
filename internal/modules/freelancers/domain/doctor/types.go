package doctor

type MedblogersID int64

type MedblogersIDs []MedblogersID

func (m MedblogersIDs) String() string {
	return ""
}

type S3Key string

func (k S3Key) String() string {
	return string(k)
}
