package engine

import (
	"crawer/fetcher"
	"fmt"
	"log"
)

func Run(seeds ...Request)  {
	var requests []Request
	for _,r:=range seeds{
		requests=append(requests,r)
	}
	for len(requests)>0{
		r:=requests[0]
		requests=requests[1:]
		
		log.Printf("<<<<<<<<<<<<<<<<<<<<Begin Fetch %s",r.Url)
		body, err := fetcher.Fetch(r.Url)
		if err!=nil{
			log.Printf("Fetcher :error Fetch url %s:%v\n",r.Url,err)
			continue
		}
		parseResult := r.ParseFunc(body)
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
		log.Printf("End Fetch %s>>>>>>>>>>>>>>>>>>>>\n\n",r.Url)
	}
}
