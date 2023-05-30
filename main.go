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
	"regexp"
)

// .通配符无法匹配换行符，而 HTML 文本中会经常出现换行符,所以用\s\S
var headerRe = regexp.MustCompile("<div class=\"small_imgposition__PYVLm\">[\\s\\S]*?<h2>([\\s\\S]*?)</h2>")

func main() {
	url := "https://www.thepaper.cn/"
	bytes, err := Fetch(url)
	if err != nil {
		fmt.Printf("read body err: %+v\n", err)
		return
	}
	matches := headerRe.FindAllSubmatch(bytes, -1)
	for _, match := range matches {
		fmt.Printf("match: %v\n", string(match[1]))
	}
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
