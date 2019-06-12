package persist

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"gopkg.in/olivere/elastic.v5"
	"spider/model"
	"testing"
)

func TestSave(t *testing.T) {
	item:=model.Profile{
			Id:1057746505,
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
	id, err := Save(item)
	if err!=nil{
		panic(err)
	}
	client, e := elastic.NewClient(elastic.SetURL(Url),elastic.SetSniff(false),)
	if e!=nil{
		panic(e)
	}
	result, err := client.Get().Index(Index).Type(Type).Id(id).Do(context.Background())
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
