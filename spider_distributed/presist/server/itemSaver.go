package main

import (
	"flag"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
	"spider/conf"
	"spider_distributed/presist"
	"spider_distributed/rpcsupport"
)
var port=flag.Int("port",0,"port to linten on")


func main() {
	flag.Parse()
	if *port==0{
		fmt.Println("please specify port")
		return
	}

	err := ServeRpc(fmt.Sprintf(":%d",*port), conf.EsConf)
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
