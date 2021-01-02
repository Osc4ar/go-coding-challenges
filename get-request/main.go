package main

import (
	"fmt"
	"net/http"
)

func contentType(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	cType := resp.Header.Get("Content-Type")
	if cType == "" {
		return cType, fmt.Errorf("Cannot get Content-Type header")
	}

	return cType, nil
}

func main() {
	contentType, err := contentType("https://www.google.com")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(contentType)
}
