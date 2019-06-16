package engine

type Parser interface {
	Parse(contents []byte, s string) ParseResult
	Serialize() (name string, args interface{})
}
type Request struct {
	Url    string
	Parser Parser
}

type ParseResult struct {
	Requests []Request
	Items    []interface{}
}
type ParseFunc func(contents []byte, s string) ParseResult
type FuncParser struct {
	parser ParseFunc
	name   string
}

func (f *FuncParser) Parse(contents []byte, name string) ParseResult {
	return f.parser(contents, name)
}

func (f *FuncParser) Serialize() (name string, args interface{}) {
	return f.name, nil
}

//工厂创建 FuncParser
func NewFuncParser(p ParseFunc, name string) *FuncParser {
	return &FuncParser{
		parser: p,
		name:   name,
	}
}

type NilParser struct {}

func (*NilParser) Parse(contents []byte, name string) ParseResult {
	return ParseResult{}
}

func (*NilParser) Serialize(_ string, _ interface{}) (string, interface{}) {
	return "NilParser", nil
}
