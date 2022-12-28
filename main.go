package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

func main() {
	response, err := http.Get("https://www.zhenai.com/zhenghun")

	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(response.Body)
	if response.StatusCode != http.StatusOK {
		fmt.Println("err: status code ", response.StatusCode)
		return
	}
	all, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	printCityList(string(all))
}

func printCityList(content string) {
	compile := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	submatch := compile.FindAllStringSubmatch(content, -1)
	for index, value := range submatch {
		fmt.Printf("%d:%s %s\n", index, value[1], value[2])
	}
}
