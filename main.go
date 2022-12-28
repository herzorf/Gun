package main

import (
	"github.com/herzorf/gun/engine"
	"github.com/herzorf/gun/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

}
