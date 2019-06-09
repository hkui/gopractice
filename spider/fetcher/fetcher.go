package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const ua  ="Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36"
var rateLimiter=time.Tick(100*time.Millisecond)

func Fetch(url string) ([]byte,error) {
	<-rateLimiter
	request,err:= http.NewRequest(http.MethodGet,url,nil)
	if err!=nil{
		panic(err)
	}
	request.Header.Add("User-Agent", ua)
	client:=http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("Redirect:",req)
			return nil

		},
	}
	resp,err:=client.Do(request)

	if err!=nil{
		return nil,err
	}
	defer resp.Body.Close()

	if resp.StatusCode!=http.StatusOK{
		return nil,fmt.Errorf("    wrong status code %d",resp.StatusCode)
	}
	bodyReader := bufio.NewReader(resp.Body)
	e:=determineEncoding(bodyReader)
	reader := transform.NewReader(bodyReader, e.NewDecoder())

	return  ioutil.ReadAll(reader)

}

func determineEncoding(r *bufio.Reader) encoding.Encoding{
	bytes, err := r.Peek(1024)
	if err!=nil{
		log.Printf("fetch err:%v",err)
		return unicode.UTF8
	}
	e,_,_ := charset.DetermineEncoding(bytes, "") //第二个返回值即为编码比如utf-8

	return e
}


