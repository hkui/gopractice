package parser

import (
	"fmt"
	"regexp"
	"spider/engine"
	"spider/model"
	"strconv"
	"strings"
)

var  reNickname  =regexp.MustCompile(`<h1 \s*class="nickName"[^<]+>([^<]+)</h1>`)

var  reId  =regexp.MustCompile(`<div class="id" [^>]+>ID：(\d+)</div>`)

var reMore=regexp.MustCompile(`<div class="des f-cl" [^>]+>([^<]+)</div>`)

func ParseProfile(contents []byte,s string) engine.ParseResult  {
	profile:=model.Profile{}
	morestr:=extractString(contents,reMore)//阿坝 | 42岁 | 大学本科 | 离异 | 163cm | 5001-8000元
	strArr:=strings.Split(morestr,"|")
	profile.Url=s
	if len(strArr)==6{
		profile.Position=strings.Trim(strArr[0]," ")
		profile.Education=strings.Trim(strArr[2]," ")
		profile.Marriage=strings.Trim(strArr[3]," ")
		profile.Income=strings.Trim(strArr[5]," ")
		ageStr:=strArr[1]
		ageStr=strings.Trim(strings.Replace(ageStr,"岁","",-1)," ")
		Age,err:=strconv.Atoi(ageStr)
		if err==nil{
			profile.Age=Age
		}else{
			fmt.Println(err)
		}
		Height,err:=strconv.Atoi(strings.Trim(strings.Replace(strArr[4],"cm","",-1)," "))
		if err==nil{
			profile.Height=Height
		}else {
			fmt.Println(err)
		}
	}
	Id,err:=strconv.Atoi(extractString(contents,reId))
	if err==nil{
		profile.Id= Id
	}
	profile.Name=extractString(contents,reNickname)
	pr:=engine.ParseResult{}
	pr.Items=append(pr.Items,profile)
	//不再向队列加reuquests了
	return pr
}

func extractString(contents []byte,re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match)>=2{
		return string(match[1])
	}
	return ""

}

type ProfileParser struct {
	userName string
}

func (p *ProfileParser) Parse(contents []byte, name string) engine.ParseResult {
	return ParseProfile(contents,p.userName)
}

func (p *ProfileParser) Serialize() (name string, args interface{}) {
	return "ParseProfile",p.userName
}


func NewProfileParser(name string) *ProfileParser  {
		return &ProfileParser{
			userName:name,
		}
}





