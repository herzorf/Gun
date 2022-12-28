package engine

type Request struct {
	Url        string
	parserFunc func([]byte) ParserResult
}

type ParserResult struct {
	Requests []Request
	item     []interface{}
}
