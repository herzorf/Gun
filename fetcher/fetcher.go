package fetcher

import (
	"fmt"
	"io"
	"net/http"
)

func Fetcher(url string) ([]byte, error) {
	response, err := http.Get(url)

	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(response.Body)

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("err status code: %d", response.StatusCode)
	}

	all, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}
	return all, nil

}
