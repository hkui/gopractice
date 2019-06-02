package main

import (
	"imooc/error/list"
	"log"
	"net/http"
	"os"
)

type appHandle func( http.ResponseWriter,  *http.Request)error

func errWrapper(handler appHandle) func( http.ResponseWriter,  *http.Request) {

	return func(writer http.ResponseWriter, request *http.Request){
		err:=handler(writer,request)
		if err!=nil{
			log.Printf("err=%s",err.Error())
			code:=http.StatusOK
			switch  {
			case os.IsNotExist(err):
				code=http.StatusNotFound
			case os.IsPermission(err):
				code=http.StatusForbidden
			default:
				code=http.StatusInternalServerError

			}
			http.Error(writer,http.StatusText(code),code)
		}
	}
}

func main()  {
	http.HandleFunc("/static/",errWrapper(list.Handle))
	err:=http.ListenAndServe(":8888",nil)
	if err!=nil{
			panic(err)
		}


}





