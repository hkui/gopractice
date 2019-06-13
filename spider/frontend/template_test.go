package frontend

import (
	"html/template"
	"os"
	"spider/frontend/model"
	"testing"
)

func TestTemplate( t *testing.T)  {
	tpl:=template.Must(template.ParseFiles("template.html"))

	page:=model.SearchResult{}

	err := tpl.Execute(os.Stdout, page)
	if err!=nil{
		panic(err)
	}
}
