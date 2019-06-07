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
	fmt.Printf("%s\n",bytes)

}

func determineEncoding(r io.Reader) encoding.Encoding{
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err!=nil{
		panic(err)
	}
	e,_,_ := charset.DetermineEncoding(bytes, "") //第二个返回值即为编码比如utf-8

	return e
}
