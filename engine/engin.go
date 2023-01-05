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

	for len(requests) > 0 {
		req := requests[0]
		requests = requests[1:]
		fmt.Printf("fetching url:%s\n", req.Url)
		content, err := fetcher.Fetcher(req.Url)
		if err != nil {
			log.Printf("fetcher err\n url: %s\n err: %v", req.Url, err)
			continue
		}
		parserResult := req.ParserFunc(content)
		requests = append(requests, parserResult.Requests...)
	}
}
