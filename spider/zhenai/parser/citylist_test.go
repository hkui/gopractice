package parser

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	//bytes, err := fetcher.Fetch("https://www.zhenai.com/zhenghun")
	bytes, err := ioutil.ReadFile("citylist_test_data.html")
	if err!=nil{
		panic(err)
	}
	list := ParseCityList(bytes)

	const resultSize  =470
	if len(list.Requests)!=resultSize{
		fmt.Printf("excepted %d,but get %d",resultSize,len(list.Requests))
	}else {
		fmt.Println("verify Success!")
	}



}
