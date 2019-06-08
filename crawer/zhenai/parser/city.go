package parser

import (
	"crawer/engine"
	"regexp"
)
//获取城市首页的用户列表
const cityRe  =`<a \s*href="(http://album.zhenai.com/u/\d+)"[^>]*>([^<]*)</a>`

func ParseCity(contents []byte)engine.ParseResult  {
	res:= regexp.MustCompile(cityRe)
	matches := res.FindAllSubmatch(contents, -1) //[][]string
	result := engine.ParseResult{}
	for _,m:=range matches{
		name:=string(m[2])
		result.Items=append(result.Items,"userlist:"+name) //昵称
		result.Requests=append(
			result.Requests,engine.Request{Url:string(m[1]),
			ParseFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c,name)
		}})
	}
	return result
}
