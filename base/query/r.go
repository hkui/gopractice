package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gpmgo/gopm/modules/log"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.zhenai.com/zhenghun/beijing")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatal("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		panic(err)
	}

	// Find the review items
	doc.Find(".m-page .paging-item").Each(func(i int, s *goquery.Selection) {
		href,b:=s.Find("a").Attr("href")
		if b{
			fmt.Println(href)
		}

	})
	html,err:=doc.Find(".m-page").Find(".paging-item").Last().Html()
	if err==nil{
		fmt.Println(html)
	}
}
