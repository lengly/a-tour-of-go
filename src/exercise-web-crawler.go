package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch 返回 URL 的 body 内容，并且将在这个页面上找到的 URL 放到一个 slice 中。
	Fetch(url string) (body string, urls []string, err error)
}

type SafeRecord struct {
	v map[string]string
	mux sync.Mutex
}

func Crawl_(url string, depth int, fetcher Fetcher, red SafeRecord, done chan bool) {
	if depth <= 0 {
		done <- true
		return
	}
	body, urls, err := fetcher.Fetch(url)
	red.mux.Lock()
	if err != nil {
		red.v[url] = "None"
		red.mux.Unlock()
		done <- true
		return
	} else {
		red.v[url] = body
		red.mux.Unlock()
	}
	for _, u := range urls {
		go Crawl_(u, depth-1, fetcher, red, done)
	}
	for i := 0; i < len(urls); i++ {
		<- done
	}
	done <- true
	return
}

// Crawl 使用 fetcher 从某个 URL 开始递归的爬取页面，直到达到最大深度。
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: 并行的抓取 URL。
	// TODO: 不重复抓取页面。
        // 下面并没有实现上面两种情况：
	red := SafeRecord{v: make(map[string]string)}
	done := make(chan bool)
	Crawl_(url, depth, fetcher, red, done)
	<- done
	for k,v := range red.v {
		if v != "None" {
			fmt.Printf("found: %s %q\n", k, v)
		} else {
			fmt.Printf("not found: %s\n", k)
		}
	}
}

func main() {
	Crawl("http://golang.org/", 4, fetcher)
}

// fakeFetcher 是返回若干结果的 Fetcher。
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher 是填充后的 fakeFetcher。
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
