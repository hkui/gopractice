package persist

import (
	"golang.org/x/net/context"
	"gopkg.in/olivere/elastic.v5"
	"log"
	"spider/conf"
)


const save = true

func ItemSaver(esConf conf.EsConfType) (chan interface{},error) {
	client, e := elastic.NewClient(elastic.SetURL(esConf["Url"]), elastic.SetSniff(false), )
	if e != nil {
		return nil,e
	}

	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("item saver got #%d,%+v", itemCount, item)
			itemCount++
			if save {
				_, err := Save(client,esConf,item)
				if err != nil {
					log.Printf("save item:%v err:%v", item, err)
				}
			}

		}

	}()
	return out,nil
}
func Save(  client *elastic.Client,esConf conf.EsConfType,item interface{}) (id string, err error) {
	response, e := client.Index().Index(esConf["Index"]).Type(esConf["Type"]).BodyJson(item).Do(context.Background())
	if e != nil {
		log.Printf("save err:%v\n",e)
		return "", e
	}

	return response.Id, nil

}
