package parser

import (
	"log"
	"regexp"
	"spider/engine"
)

//获取城市首页的用户列表
var cityRe = regexp.MustCompile(`<a \s*href="(http://album.zhenai.com/u/\d+)"[^>]*>([^<]*)</a>`)

var nextCityRe = regexp.MustCompile(`<li class="paging-item">\s*<a href="([^"]+)">下一页</a>`)

func ParseCity(contents []byte) engine.ParseResult {

	matches := cityRe.FindAllSubmatch(contents, -1) //[][]string
	result := engine.ParseResult{}
	for _, m := range matches {
			url:=string(m[1])
		//result.Items = append(result.Items, "userlist:"+name) //昵称
		result.Requests = append(
			result.Requests, engine.Request{
				Url: string(m[1]),
				/*ParseFunc: func(c []byte) engine.ParseResult {
					return ParseProfile(c, name)
				},*/
				ParseFunc:profileParser(url),//可以直接用string(m[1]) ，函数调用本身就是拷贝
			})
	}

	mnext := nextCityRe.FindAllSubmatch(contents, 1)
	for _, m1 := range mnext {
		log.Printf("next=%s\n",m1[1])
		result.Requests = append(
			result.Requests, engine.Request{
				Url:       string(m1[1]),
				ParseFunc: ParseCity,
			})

	}

	return result
}
