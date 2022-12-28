package parser

import (
	"github.com/herzorf/gun/engine"
	"regexp"
)

const cityListReg = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(content []byte) engine.ParserResult {
	compile := regexp.MustCompile(cityListReg)
	submatch := compile.FindAllStringSubmatch(string(content), -1)
	result := engine.ParserResult{}
	for _, value := range submatch {
		result.Item = append(result.Item, value[2])
		result.Requests = append(result.Requests, engine.Request{
			Url:        value[1],
			ParserFunc: engine.NilParser,
		})
	}
	return result
}
