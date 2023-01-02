package main

import (
	"github.com/herzorf/gun/engine"
	"github.com/herzorf/gun/miyoushe/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "https://www.miyoushe.com/ys/postRanking?forum_id=49&date=49-202212",
		ParserFunc: parser.MHYArticleListParse,
	})

}
