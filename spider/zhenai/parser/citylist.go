package parser

import (
	"regexp"
	"spider/engine"
)

var cityListRe  =regexp.MustCompile( `<a \s*href="(http://www.zhenai.com/zhenghun/\w+)"[^>]*>([^<]+)</a>`)
//<a href="http://www.zhenai.com/zhenghun/aba" data-v-5e16505f="">阿坝</a>
//获取城市列表
func ParseCityList(contents []byte)engine.ParseResult  {
	limit:=10

	matches := cityListRe.FindAllSubmatch(contents,-1) //[][]string
	result := engine.ParseResult{}
	for _,m:=range matches{
		result.Items=append(result.Items,"citylist:"+string(m[2])) //城市中文名
		result.Requests=append(result.Requests,engine.Request{Url:string(m[1]),ParseFunc:ParseCity})
		limit--
		if limit<=0{
			break
		}
	}
	return result
}
