package collect

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

type Fetcher interface {
	Get(url string) ([]byte, error)
}

type BaseFetcher struct {
}

func (b *BaseFetcher) Get(url string) ([]byte, error) {
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

type BrowserFetcher struct {
}

func (b *BrowserFetcher) Get(url string) ([]byte, error) {
	c := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "http.NewRequest err")
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36")
	rsp, err := c.Do(req)
	if err != nil {
		return nil, errors.Wrapf(err, "http.Do err")
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
