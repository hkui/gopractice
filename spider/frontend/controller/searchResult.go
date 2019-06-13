package controller

import (
	"fmt"
	"gopkg.in/olivere/elastic.v5"
	"net/http"
	"spider/frontend/views"
	"strconv"
	"strings"
)

type SearchResultHandler struct {
	view views.SearchResultView
	client *elastic.Client
}


func (h SearchResultHandler) ServeHTTP(w http.ResponseWriter,req *http.Request)  {
	q:=strings.TrimSpace(req.FormValue("q"))

	from, err := strconv.Atoi(req.FormValue("from"))
	if err!=nil{
		from=0
	}
	fmt.Fprintf(w,"q=%s,from=%d\n",q,from)
}
