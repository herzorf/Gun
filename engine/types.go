package engine

import (
	"fmt"
)

type Request struct {
	Url        string
	Data       ArticleCard
	ParserFunc func(Request, []byte) ParserResult
}

type ParserResult struct {
	Requests []Request
}

func NilParser(req Request, content []byte) ParserResult {
	fmt.Println(req)
	return ParserResult{}
}

type ArticleCard struct {
	Username string
	Date     string
	Title    string
	Topics   []string
	Link     string
	Info     interface{}
	DataSrc  []string
}
