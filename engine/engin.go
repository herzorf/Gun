package engine

import (
	"fmt"
	"github.com/herzorf/gun/fetcher"
	"log"
)

func Run(seed ...Request) {
	var requests []Request

	for _, req := range seed {
		requests = append(requests, req)
	}
	var count = 5
	for len(requests) > 0 {
		if count < 0 {
			break
		}
		count--
		req := requests[0]
		requests = requests[1:]
		fmt.Printf("fetching url:%s\n", req.Url)
		content, err := fetcher.Fetcher(req.Url)
		if err != nil {
			log.Printf("fetcher err\n url: %s\n err: %v", req.Url, err)
			continue
		}
		parserResult := req.ParserFunc(req, content)
		requests = append(requests, parserResult.Requests...)
	}
}
