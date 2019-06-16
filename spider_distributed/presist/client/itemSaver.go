package client

import (
	"log"
	"spider_distributed/rpcsupport"
)

func ItemSaver(host string) (chan interface{},error) {

	client, err := rpcsupport.NewClient(host)
	if err!=nil{
		panic(err)
	}
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("item saver got #%d,%+v", itemCount, item)
			itemCount++
			result:=""
			e:= client.Call("ItemSaverService.Save", item, &result)
			if e!=nil || result!="ok"{
				log.Printf("err:%v,item:%v\n",e,item)
			}
		}

	}()
	return out,nil
}

