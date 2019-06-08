package parser

import (
	"fmt"
	"spider/fetcher"
	"testing"
)

func TestParseProfile(t *testing.T) {
	bytes, e := fetcher.Fetch("http://album.zhenai.com/u/1342436087")
	if e!=nil{
		panic(e)
	}

	profile := ParseProfile(bytes,"hk")
	fmt.Printf("%v",profile.Items)

}
