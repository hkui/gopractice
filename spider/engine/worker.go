package engine

import (
	"log"
	"spider/fetcher"
)

//从request获取parseRequest
func Worker(r Request)(ParseResult,error)  {
	body, err := fetcher.Fetch(r.Url)
	if err!=nil{
		log.Printf("Fetcher :error Fetch url %s:%v\n",r.Url,err)
		return ParseResult{},err
	}
	parseResult := r.Parser.Parse(body,"")
	return parseResult,nil
}
