package proxy

import (
	"github.com/pkg/errors"
	"net/http"
	"net/url"
	"sync/atomic"
)

type ProxyFunc func(*http.Request) (*url.URL, error)

type roundRobinSwitcher struct {
	urls  []*url.URL
	index uint32
}

func (r *roundRobinSwitcher) GetProxy(*http.Request) (*url.URL, error) {
	r.index = atomic.AddUint32(&r.index, 1) % uint32(len(r.urls))
	return r.urls[r.index], nil
}

func RoundRobinProxySwitcher(urls ...string) (ProxyFunc, error) {
	if len(urls) < 1 {
		return nil, errors.New("invalid urls")
	}
	urlList := make([]*url.URL, 0, len(urls))
	for _, u := range urls {
		parsedU, err := url.Parse(u)
		if err != nil {
			return nil, errors.Wrapf(err, "url.Parse err")
		}
		urlList = append(urlList, parsedU)
	}
	switcher := &roundRobinSwitcher{
		urls:  urlList,
		index: 0,
	}
	return switcher.GetProxy, nil
}
