package persist

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"gopkg.in/olivere/elastic.v5"
	"spider/conf"
	"spider/model"
	"testing"
)

func TestSave(t *testing.T) {
	item:=model.Profile{
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
	esConf:=conf.EsConf
	client, e := elastic.NewClient(elastic.SetURL(esConf["Url"]), elastic.SetSniff(false), )
	if e != nil {
		panic(e)
	}
	id, err := Save(client,esConf,item)
	if err!=nil{
		panic(err)
	}

	result, err := client.Get().Index(esConf["Index"]).Type(esConf["Type"]).Id(id).Do(context.Background())
	if err!=nil{
		panic(err)
	}
	fmt.Printf("%s\n",result.Source)
	var actual model.Profile

	err = json.Unmarshal([]byte(*result.Source), &actual)
	if err!=nil{
		panic(err)
	}
	fmt.Printf("%v",actual)

}
