package main

import (
	"net/http"
	"spider/conf"
	"spider/frontend/controller"
)

func main()  {
	http.Handle("/",http.FileServer(http.Dir("spider/frontend/views")))

	http.Handle("/search",
		controller.CreateSearchResulthandler(
			"spider/frontend/views/template.html",conf.EsConf,20)) //必须实现ServeHTTP
	err := http.ListenAndServe(":8888",nil)
	if err!=nil{
		panic(err)
	}
}
