package mock

type Mretriever struct {
	Contents string

}

func (r Mretriever) Get(url string) string {
	return  r.Contents
}

