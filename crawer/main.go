package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	resp, err := http.Get("https://www.zhenai.com/zhenghun")
	if err!=nil{
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode!=http.StatusOK{
		fmt.Println("Err:status code:",resp.StatusCode)
		return
	}
	e:=determineEncoding(resp.Body)
	reader := transform.NewReader(resp.Body, e.NewDecoder())

	bytes, err := ioutil.ReadAll(reader)
	if err!=nil{
		panic(err)
	}
	printCityList(bytes)

}

func determineEncoding(r io.Reader) encoding.Encoding{
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err!=nil{
		panic(err)
	}
	e,_,_ := charset.DetermineEncoding(bytes, "") //第二个返回值即为编码比如utf-8

	return e
}
func printCityList(contents []byte)  {
	//<a href="http://www.zhenai.com/zhenghun/aba" data-v-5e16505f="">阿坝</a>
	res:= regexp.MustCompile(`<a \s*href="(http://www.zhenai.com/zhenghun/\w+)"[^>]*>([^<]+)</a>`)
	matches := res.FindAllSubmatch(contents, -1) //[][]string
	for _,m:=range matches{
		fmt.Printf("url=>%s        name=>%s\n",m[1],m[2])

		fmt.Println()

	}

}
