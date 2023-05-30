package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "https://www.thepaper.cn/"
	rsp, err := http.Get(url)
	if err != nil {
		fmt.Printf("fetch url err: %+v\n", err)
		return
	}
	if rsp.StatusCode != http.StatusOK {
		fmt.Printf("rsp status not 200\n")
		return
	}
	bytes, err := io.ReadAll(rsp.Body)
	if err != nil {
		fmt.Printf("read body err: %+v\n", err)
		return
	}
	fmt.Println("body: ", string(bytes))
}
