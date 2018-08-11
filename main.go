package main

import (
	"fmt"

	"github.com/go-resty/resty"
)

func main() {
	resp, err := resty.R().Get("http://httpbin.org/get")

	fmt.Println(resp)
	fmt.Println(err)
	fmt.Println("hello world")
}
