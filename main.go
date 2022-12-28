package main

import (
	"fmt"
	"io"
	"net/http"
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
	fmt.Println(string(all))
}
