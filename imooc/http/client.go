package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	request,err:= http.NewRequest(http.MethodGet,"http://www.imooc.com/",nil)
	if err!=nil{
		panic(err)
	}
	request.Header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 5.0; SM-G900P Build/LRX21T) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Mobile Safari/537.36")
	client:=http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("REdirect:",req)
			return nil

		},
	}

	resp,err:=client.Do(request)
	if err!=nil{
		panic(err)
	}
	defer resp.Body.Close()
	s,err:=httputil.DumpResponse(resp,true)
	if err!=nil{
		panic(err)
	}
	fmt.Printf("s=%s\n",s)
}
