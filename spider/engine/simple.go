package engine

import (
	"fmt"
	"log"
)
type SimpleEngine struct {

}

func (e SimpleEngine )Run(seeds ...Request)  {
	var requests []Request
	for _,r:=range seeds{
		requests=append(requests,r)
	}
	for len(requests)>0{
		r:=requests[0]
		requests=requests[1:]
		
		parseResult,err:=worker(r)
		if err!=nil{
			continue
		}

		RequestsLen:=len(parseResult.Requests)

		if(RequestsLen>0){
			requests= append(requests, parseResult.Requests ...)
		}
		for k,item:=range parseResult.Items{
			logger:=""
			logger +=fmt.Sprintf("%v    ",item)
			if k<RequestsLen{
				logger+=fmt.Sprintf(" %s",parseResult.Requests[k].Url)
			}
			log.Println(logger)
		}
	}
}


