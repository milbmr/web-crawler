package main

import (
	"fmt"
	"sync"

	"github.com/milbmr/web-crawler/crawler"
)

func main() {
	var w sync.WaitGroup
	url := make(chan string)
	done := make(chan interface{})
	defer close(done)
	w.Add(2)
	defer w.Done()
	go crawler.Crawl(done, 2, "http://golang.org/", url, &w, fetcher)

	go func() {
		defer w.Done()
		for u := range url {
			fmt.Println(u)
		}
	}()

	w.Wait()
  fmt.Println("exe")
	close(url)
}

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
