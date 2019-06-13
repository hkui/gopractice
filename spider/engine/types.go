package engine

type ParseFunc func([]byte) ParseResult
type Item []interface{}
type Request struct {
	Url string
	ParseFunc ParseFunc
}

type ParseResult struct {
	Requests []Request
	Items Item
}

func NilParser(contents []byte) ParseResult {
	return ParseResult{}
}
