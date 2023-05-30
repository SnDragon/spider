package main

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
	"net/http"
)

func main() {
	url := "https://www.thepaper.cn/"
	bytes, err := Fetch(url)
	if err != nil {
		fmt.Printf("read body err: %+v\n", err)
		return
	}
	fmt.Printf("rsp body: %v\n", string(bytes))
}

func Fetch(url string) ([]byte, error) {
	rsp, err := http.Get(url)
	if err != nil {
		return nil, errors.Wrapf(err, "fetch url")
	}
	if rsp.StatusCode != http.StatusOK {
		fmt.Printf("rsp status not 200\n")
		return nil, errors.Errorf("rsp status: %v not 200", rsp.StatusCode)
	}
	reader := bufio.NewReader(rsp.Body)
	e := DetermineEncoding(reader)
	utf8Reader := transform.NewReader(reader, e.NewDecoder())
	return io.ReadAll(utf8Reader)
}

func DetermineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		fmt.Printf("DetermineEncoding Peek err: %+v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
