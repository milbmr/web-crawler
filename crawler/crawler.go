package crawler

import (
	"io"
	"log"
	"net/http"
)

func Crawl(done <-chan interface{}, depth int, urlStream string, out chan string) {
	if depth <= 0 {
		return
	}

	visited := make(map[string]bool)
	if visited[urlStream] {
		return
	}

	html := getHtml(done, urlStream, visited)
	u := GenerateUrls(done, html)
	go Crawl(done, depth-1, <-u, out)

	for {
		select {
		case <-done:
			return
		}
	}
}

func getHtml(done <-chan interface{}, urlStream string, visited map[string]bool) <-chan []byte {
	visited[urlStream] = true
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
