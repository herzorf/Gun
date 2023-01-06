package engine

import "fmt"

type Request struct {
	Url        string
	Data       interface{}
	ParserFunc func(Request, []byte) ParserResult
}

type ParserResult struct {
	Requests []Request
}

func NilParser(req Request, content []byte) ParserResult {
	fmt.Println(req)
	return ParserResult{}
}
