package main

import (
	"crawer/engine"
	"crawer/zhenai/parser"
)

const seedUrl ="https://www.zhenai.com/zhenghun"
func main() {
	Requests:=engine.Request{Url:seedUrl,ParseFunc:parser.ParseCityList}

	engine.Run(Requests)
}



