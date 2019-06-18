package worker

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
	"spider/engine"
	"spider/zhenai/parser"
	"spider_distributed/config"
)

type SerializedParser struct {
	Name string
	Args interface{}
}
type Request struct {
	Url string
	Parser SerializedParser
}

func SerializeRequest(r engine.Request) Request {
	name,args:=r.Parser.Serialize()
	return Request{
		Url:r.Url,
		Parser:SerializedParser{
			Name:name,
			Args:args,
		},
	}
}
func SerializeResult(r engine.ParseResult)ParseResult  {
	result:=ParseResult{
		Items:r.Items,
	}
	for _,req:=range r.Requests{
		result.Requests=append(result.Requests,SerializeRequest(req))
	}
	return result
}

func DeserializeRequest(r Request)(engine.Request,error){
	parser,err:=DeserializeParser(r.Parser)
	if err!=nil{
		return engine.Request{},err
	}
	return engine.Request{
		Url:r.Url,
		Parser:parser,
	},nil
}
func DeserializeResult(r ParseResult) engine.ParseResult{
	result:=engine.ParseResult{
		Items:r.Items,
	}
	for _,req:=range r.Requests{
		engineReq,err:=DeserializeRequest(req)
		if err!=nil{
			log.Printf("error DeserializeRequest :%v",err)
		}
		result.Requests=append(result.Requests,engineReq)
	}
	return result
}
func DeserializeParser( p SerializedParser) (engine.Parser,error) {
	switch p.Name{
	case config.ParseCityList:
		return engine.NewFuncParser(parser.ParseCityList,config.ParseCityList),nil
	case config.ParseCity:
		return engine.NewFuncParser(parser.ParseCity,config.ParseCity),nil
	case config.NilParser:
		return engine.NilParser{},nil
	case config.ParseProfile:
		if userName,ok:=p.Args.(string);ok{
			return parser.NewProfileParser(userName),nil
		}else{
			return nil,fmt.Errorf("invalid "+"args:%v",p.Args)
		}
	default:
		return nil,errors.Errorf("unknow parser name:%s\n",p.Name)
	}
}
type ParseResult struct {
	Items []interface{}
	Requests []Request
}
















