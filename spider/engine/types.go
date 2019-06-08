package engine

type Request struct {
	Url string
	ParseFunc func([]byte) ParseResult
}
type ParseResult struct {
	Requests []Request
	Items []interface{}
}

func NilParser(contents []byte) ParseResult {
	return ParseResult{}
}
