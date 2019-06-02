package mock

import "fmt"

type Mretriever struct {
	Contents string

}

func (r *Mretriever) String() string {
	return fmt.Sprintf("mretriver [contents:%s]",r.Contents)
}

func (r *Mretriever) Post(url string, form map[string]string) string {
	r.Contents=form["contents"]
	return "ok"

}

func (r *Mretriever) Get(url string) string {
	return  r.Contents
}

