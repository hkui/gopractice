package worker

import (
	"spider/engine"
)
//要注册的服务
type CrawlService struct {

}

func (CrawlService)Process(req Request,result *ParseResult)error  {
	engineReq,err:=DeserializeRequest(req)
	if err!=nil{
		return err
	}
	engineResult,err:=engine.Worker(engineReq)
	if err!=nil{
		return err
	}
	*result=SerializeResult(engineResult)
	return nil
}
//必须返回error类型嘛
func (CrawlService)Test(arg int ,res *int) error {
	*res=arg
	return nil
}




















