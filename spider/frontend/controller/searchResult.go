package controller

import (
	"fmt"
	"golang.org/x/net/context"
	"gopkg.in/olivere/elastic.v5"
	"log"
	"net/http"
	"reflect"
	"spider/conf"
	"spider/frontend/model"
	"spider/frontend/views"
	model2 "spider/model"
	"strconv"
	"strings"
)

// es的处理
type EsHandle struct {
	esConf conf.EsConfType
	client *elastic.Client
}
//模板处理
type ViewResult struct {
	view views.SearchResultView
}
type SearchResultHandler struct {
	ViewResult
	EsHandle
	size int
}

func (h SearchResultHandler) ServeHTTP(w http.ResponseWriter,req *http.Request)  {
	q:=strings.TrimSpace(req.FormValue("q"))
	from, err := strconv.Atoi(req.FormValue("from"))
	if err!=nil{
		from=0
	}
	var page model.SearchResult

	page,err=h.getSearchResult(q,from,h.size)
	log.Printf("q=%s,from=%d\n",q,from)
	if err!=nil{
		http.Error(w,err.Error(),http.StatusBadRequest)
	}
	page.Query=q
	err = h.view.Render(w, page)
	if err!=nil{
		http.Error(w,err.Error(),http.StatusBadRequest)
	}

	fmt.Fprintf(w,"q=%s,from=%d\n",q,from)
}
//得到es处理的handle,模板的处理handle
func CreateSearchResulthandler(template string,esConf conf.EsConfType,size int) SearchResultHandler{
	client, e := elastic.NewClient(elastic.SetURL(esConf["Url"]), elastic.SetSniff(false), )
	if e!=nil{
		panic(e)
	}
	viewResult:=views.CreateSearchResultView(template)
	var SearchResultHandler SearchResultHandler
	SearchResultHandler.ViewResult.view=viewResult
	SearchResultHandler.EsHandle.client=client
	SearchResultHandler.size=size
	return SearchResultHandler

}

//Es根据查询条件获取结果
func ( Es EsHandle)getSearchResult(q string,from int,size int)(model.SearchResult,error){
	var result model.SearchResult
	esConf:=Es.esConf
	var resp *elastic.SearchResult
	var err error
	if len(q)>0{
		query:=elastic.NewQueryStringQuery(q)
		resp,err=Es.client.Search(esConf["Index"]).Type(esConf["Type"]).Query(query).From(from).Size(size).Do(context.Background())
	}else{
		resp,err=Es.client.Search(esConf["Index"]).Type(esConf["Type"]).From(from).Size(size).Do(context.Background())
	}

	if err!=nil{
		return result,err
	}
	result.Hits=resp.TotalHits()
	result.Start=from
	result.Items=resp.Each(reflect.TypeOf(model2.Profile{}))


	return result,nil

}

