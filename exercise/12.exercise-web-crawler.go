/**
 * 练习：Web 爬虫
 * -------------
在这个练习中，将会使用 Go 的并发特性来并行执行 web 爬虫。

修改 Crawl 函数来并行的抓取 URLs，并且保证不重复。

提示：你可以用一个 map 来缓存已经获取的 URL，但是需要注意 map 本身并不是并发安全的！
*/
package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch 返回 URL 的 body 内容，并且将在这个页面上找到的 URL 放到一个 slice 中。
	Fetch(url string) (body string, urls []string, err error)
}

var m = make(map[string]int)
var mux sync.Mutex	// 互斥锁

// Crawl 使用 fetcher 从某个 URL 开始递归的爬取页面，直到达到最大深度。
func Crawl(url string, depth int, fetcher Fetcher) {
	if depth <= 0 {
		return
	}

	mux.Lock()
	if m[url] == 1 { // 检查url是否已经被 fetch 过，fetch 过则跳过
		mux.Unlock()
		return
	}

	body, urls, err := fetcher.Fetch(url)
	m[url] = 1
	mux.Unlock()

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)

	for _, u := range urls {
		go Crawl(u, depth-1, fetcher)  // 并行，根据url的数量开启多个goroutine
	}
	return
}

func main() {
	Crawl("http://golang.org/", 4, fetcher) // 并行，开启一个goroutine
}

// fakeFetcher 是返回若干结果的 Fetcher。
type fakeResult struct {
	body string
	urls []string
}

type fakeFetcher map[string]*fakeResult

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