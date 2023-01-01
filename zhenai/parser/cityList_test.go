package parser

import (
	"fmt"
	"os"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := os.ReadFile("cityList_test_data.html")
	if err != nil {
		panic(err)
	}
	list := CityList(contents)
	fmt.Println(len(list.Requests))

	const resultSize = 470
	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba", "http://www.zhenai.com/zhenghun/akesu", "http://www.zhenai.com/zhenghun/alashanmeng",
	}
	expectedCitys := []string{
		"阿坝", "阿克苏", "阿拉善盟",
	}
	if len(list.Requests) != resultSize {
		t.Errorf("list should have %d"+"requests; but had %d", resultSize, len(list.Requests))
	}
	for i, url := range expectedUrls {
		if list.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s; but was %s", i, url, list.Requests[i].Url)
		}
	}
	if len(list.Item) != resultSize {
		t.Errorf("list should have %d"+"requests; but had %d", resultSize, len(list.Item))
	}
	for i, item := range expectedCitys {
		if list.Item[i] != item {
			t.Errorf("expected url #%d: %s; but was %s", i, item, list.Item[i])
		}
	}
}
