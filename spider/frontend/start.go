package main

import (
	"net/http"
	"spider/frontend/controller"
)

func main()  {
	http.Handle("/search",controller.SearchResultHandler{}) //必须实现ServeHTTP
	err := http.ListenAndServe(":8888",nil)
	if err!=nil{
		panic(err)
	}

}
