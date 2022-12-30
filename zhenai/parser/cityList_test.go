package parser

import (
	"github.com/herzorf/gun/fetcher"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := fetcher.Fetcher("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	CityList(contents)
}
