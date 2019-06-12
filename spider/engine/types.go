package engine

type ParseFunc func([]byte) ParseResult
type Request struct {
	Url string
	ParseFunc ParseFunc
}
type ParseResult struct {
	Requests []Request
	Items []interface{}
}

func NilParser(contents []byte) ParseResult {
	return ParseResult{}
}
