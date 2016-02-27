/* A Tour of Go (73): Exercise: Web Crawler. */
package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

type CrawlHistory struct {
	urls map[string]bool
	sync.Mutex
}

func Crawl(url string, depth int, fetcher Fetcher) {
	h := CrawlHistory{urls: make(map[string]bool)}
	wg := sync.WaitGroup{}

	var crawl func(url string, depth int)
	crawl = func(url string, depth int) {
		defer wg.Done()

		h.Lock()
		if h.urls[url] {
			h.Unlock()
			return
		} else {
			h.urls[url] = true
		}
		h.Unlock()

		body, urls, err := fetcher.Fetch(url)
		if err == nil {
			fmt.Printf("found: %s %q\n", url, body)
		} else {
			fmt.Println(err)
		}

		if depth-1 > 0 {
			for _, u := range urls {
				wg.Add(1)
				go crawl(u, depth-1)
			}
		}
	}

	wg.Add(1)
	go crawl(url, depth)
	wg.Wait()
}

func main() {
	Crawl("http://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
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

// fetcher is a populated fakeFetcher.
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
