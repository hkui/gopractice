package main

import (
	"gopkg.in/olivere/elastic.v5"
	"spider/conf"
	"spider_distributed/presist"
	"spider_distributed/rpcsupport"
)

func main() {
	err := ServeRpc(":1234", conf.EsConf)
	if err!=nil{
		panic(err)
	}
}

func  ServeRpc(host string,esConf conf.EsConfType)error  {
	client, e := elastic.NewClient(elastic.SetURL(esConf["Url"]), elastic.SetSniff(false), )
	if e != nil {
		panic(e)
	}

	e= rpcsupport.ServeRpc(host, &presist.ItemSaverService{Client: client, Esconf: esConf})
	if e!=nil{
		panic(e)
	}
	return e

}
