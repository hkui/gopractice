package parser

import (
	"log"
	"regexp"
	"spider/engine"
)

//获取城市首页的用户列表
var cityRe = regexp.MustCompile(`<a \s*href="(http://album.zhenai.com/u/\d+)"[^>]*>([^<]*)</a>`)

var nextCityRe = regexp.MustCompile(`<li class="paging-item">\s*<a href="([^"]+)">下一页</a>`)

func ParseCity(contents []byte,name string) engine.ParseResult {

	matches := cityRe.FindAllSubmatch(contents, -1) //[][]string
	result := engine.ParseResult{}
	for _, m := range matches {
			url:=string(m[1])
		//result.Items = append(result.Items, "userlist:"+name) //昵称
		result.Requests = append(
			result.Requests, engine.Request{
				Url: string(m[1]),
				Parser:NewProfileParser(url),
			})
	}
	//匹配下一页
	mnext := nextCityRe.FindAllSubmatch(contents, 1)
	for _, m1 := range mnext {
		log.Printf("next=%s\n",m1[1])
		result.Requests = append(
			result.Requests, engine.Request{
				Url:       string(m1[1]),
				Parser: engine.NewFuncParser(ParseCity,"ParseCity"),
			})

	}

	return result
}
