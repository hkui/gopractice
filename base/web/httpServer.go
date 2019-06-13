package main

import (
	"fmt"
	"github.com/gpmgo/gopm/modules/log"
	"net/http"
	"runtime"
	"strings"
)

func say(w http.ResponseWriter,r *http.Request)  {
	fmt.Printf("-----------------------------------\n")
	r.ParseForm()
	header:="<h2>headers</h2>"
	for k,v:=range r.Header{
		header+=fmt.Sprintf("%s:%s<br>",k,strings.Join(v,""))
	}
	header+="<hr>"


	fmt.Printf("url=%+v\n",r.URL.Path)
	fmt.Printf("method=%s\n",r.Method)
	html:=header
	html+="method:"+r.Method+"<br>";
	html+="url:="+r.URL.Path+"<br>"
	for k, v := range r.Form {
		html+=fmt.Sprintf("%s=%v<br>",k,strings.Join(v,""))
		fmt.Println(html)
	}
	w.Header().Set("Content-type","text/html; charset=UTF-8")
	w.Header().Set("X-Powered-By:",runtime.Version())
	fmt.Fprintln(w,html)

}

func main() {
	http.HandleFunc("/",say)
	err := http.ListenAndServe(":9090", nil)
	if err!=nil{
		log.Fatal("ListenAndServe： ",err)
	}
}





