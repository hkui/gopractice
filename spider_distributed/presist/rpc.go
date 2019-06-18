package presist

import (
	"gopkg.in/olivere/elastic.v5"
	"log"
	"spider/conf"
	"spider/persist"
)
//要注册进去的rpc服务
type ItemSaverService struct {
	Client *elastic.Client
	Esconf conf.EsConfType
}
var count=0

func (s *ItemSaverService)Save(item interface{},result *string)error  {
	_, err := persist.Save(s.Client, s.Esconf, item)
	if err==nil{
		count++
		*result="ok"
		log.Printf("rpc save %d, item:%v\n ",count,item)
	}
	return  err
}

