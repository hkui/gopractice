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

type SearchResultHandler struct {
	view views.SearchResultView
	esConf conf.EsConfType
	client *elastic.Client
}

func (h SearchResultHandler) ServeHTTP(w http.ResponseWriter,req *http.Request)  {
	q:=strings.TrimSpace(req.FormValue("q"))

	from, err := strconv.Atoi(req.FormValue("from"))
	if err!=nil{
		from=0
	}
	var page model.SearchResult

	page,err=h.getSearchResult(q,from)
	log.Printf("q=%s,from=%d\n",q,from)
	if err!=nil{
		http.Error(w,err.Error(),http.StatusBadRequest)
	}
	err = h.view.Render(w, page)
	if err!=nil{
		http.Error(w,err.Error(),http.StatusBadRequest)
	}

	fmt.Fprintf(w,"q=%s,from=%d\n",q,from)
}
func CreateSearchResulthandler(template string,esConf conf.EsConfType) SearchResultHandler{
	client, e := elastic.NewClient(elastic.SetURL(esConf["Url"]), elastic.SetSniff(false), )
	if e!=nil{
		panic(e)
	}
	return SearchResultHandler{
		view:views.CreateSearchResultView(template),
		client:client,
	}

}
func ( sh SearchResultHandler)getSearchResult(q string,from int)(model.SearchResult,error){
	var result model.SearchResult
	esConf:=sh.esConf
	//query:=elastic.NewQueryStringQuery(q)
	resp,err:=sh.client.Search(esConf["Index"]).Type(esConf["Type"])/*.Query(query)*/.From(from).Do(context.Background())
	if err!=nil{
		return result,err
	}
	result.Hits=resp.TotalHits()
	result.Start=from
	result.Items=resp.Each(reflect.TypeOf(model2.Profile{}))
	return result,nil

}

