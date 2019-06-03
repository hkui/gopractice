package main

import (
	"imooc/error/list"
	"log"
	"net/http"
	"os"
)

func errWrapper(handler list.AppHandle) func( http.ResponseWriter,  *http.Request) {

	return func(writer http.ResponseWriter, request *http.Request){

		defer func() {
			if r:=recover();r!=nil{
				log.Printf("panic:%v",r)
				http.Error(writer,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)

			}
		}()
		
		err:=handler(writer,request)
		if err!=nil{
			if userErr,ok:=err.(userError);ok{
				http.Error(writer,userErr.Message(),http.StatusBadRequest)
				return
			}


			code:=http.StatusOK
			log.Printf("err=%s,code=%d\n",err.Error(),code)
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

type userError interface {
	error
	Message() string
}

func main()  {
	http.HandleFunc("/",errWrapper(list.Handle))
	err:=http.ListenAndServe(":8888",nil)
	if err!=nil{
			panic(err)
		}


}





