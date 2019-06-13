package views

import (
	"os"
	"spider/frontend/model"
	model2 "spider/model"
	"testing"
)

func TestCreateSearchResultView( t *testing.T)  {
	view:=CreateSearchResultView("template.html")
	page:=model.SearchResult{}

	item:=model2.Profile{
		Id:1057746505,
		Url:"http://album.zhenai.com/u/1057746505",
		Name:"何处不相逢",
		Gender:"",
		Age:37 ,
		Height:158 ,
		Weight:0 ,
		Income:"12001-20000元",
		Marriage:"离异",
		Education:"硕士",
		Position:"北京",
	}
	for i:=0;i<10;i++{
		page.Items=append(page.Items,item)
	}
	page.Hits=len(page.Items)

	//err := tpl.Execute(os.Stdout, page)//直接输出
	file, err := os.Create("template_test.html")
	err=view.Render(file,page)
	if err!=nil{
		panic(err)
	}

}
