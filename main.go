package main

import (
	"fmt"
	"github.com/herzorf/gun/fetcher"
	"regexp"
)

func main() {
	content, err := fetcher.Fetcher("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	printCityList(string(content))
}

func printCityList(content string) {
	compile := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	submatch := compile.FindAllStringSubmatch(content, -1)
	for index, value := range submatch {
		fmt.Printf("%d:%s %s\n", index, value[1], value[2])
	}
}
