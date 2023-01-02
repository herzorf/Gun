package engine

type Request struct {
	Url        string
	Data       interface{}
	ParserFunc func([]byte) ParserResult
}

type ParserResult struct {
	Requests []Request
}

func NilParser([]byte) ParserResult {
	return ParserResult{}
}
