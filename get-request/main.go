package main

import (
	"fmt"
	"net/http"
)

func contentType(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("Cannot get content type of %s", url)
	}
	defer resp.Body.Close()

	return resp.Header.Get("Content-Type"), nil
}

func main() {
	contentType, err := contentType("https://www.google.com")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(contentType)
}
