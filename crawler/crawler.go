package crawler

import (
	"io"
	"log"
	"net/http"
)

func Crawler(done <-chan interface{}, baseUrl string, depth int, urlStream <-chan string) {
	if depth <= 0 {
		return
	}

  urls := make(chan string)
	html := getHtml(done, urlStream)
	Parser(done, html)
}

func getHtml(done <-chan interface{}, urlStream <-chan string) <-chan []byte {
	html := make(chan []byte)
	go func() {
		defer close(html)
		res, err := http.Get(<-urlStream)
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
