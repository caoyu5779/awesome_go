package main

import (
	"selfLearning/crawler/engine"
	"selfLearning/crawler/zhenai/parser"
)

func main() {

	engine.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun",
		ParserFunc:parser.ParseCityList,
	})
}

