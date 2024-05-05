package crawler

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

func Crawl(done <-chan interface{}, depth int, urlStream string, out chan string, w *sync.WaitGroup, fetcher Fetcher) {
	defer w.Done()
	if depth <= 0 {
		return
	}

	visited := make(map[string]bool)
	var mux sync.Mutex

	mux.Lock()
	if ok := visited[urlStream]; !ok {
		visited[urlStream] = true
		mux.Unlock()
	} else {
		return
	}

	_, urls, err := fetcher.Fetch(urlStream)
	if err != nil {
		fmt.Fprintln(os.Stdout, "err fetching", err)
		return
	}
	u := make(chan string)
	go func() {
		defer close(u)
		for _, ur := range urls {
			select {
			case <-done:
				return
			case u <- ur:
			}
		}
	}()
	// html := getHtml(done, urlStream)
	// u := GenerateUrls(done, html)

	// for {
	// 	select {
	// 	case <-done:
	// 		return
	// 	case val := <-u:
	// 		w.Add(1)
	// 		go Crawl(done, depth-1, val, out, w, fetcher)
	// 		out <- val
	// 	}
	// }
	for url := range u {
    w.Add(1)
		go Crawl(done, depth-1, url, out, w, fetcher)
		out <- url
	}
}

func getHtml(done <-chan interface{}, urlStream string) <-chan []byte {
	html := make(chan []byte)
	go func() {
		defer close(html)
		res, err := http.Get(urlStream)
		if err != nil {
			log.Fatal(err)
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		res.Body.Close()
		select {
		case <-done:
			return
		case html <- body:
		}
	}()

	return html
}
