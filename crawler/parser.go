package crawler

import (
	"bytes"
	"log"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func parse(page []byte) []string {
	var urls []string
	doc, err := html.Parse(bytes.NewReader(page))
	if err != nil {
		log.Fatal(err)
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" && isUrl(a.Val) {
					urls = append(urls, a.Val)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return urls
}

func GenerateUrls(done <-chan interface{}, page <-chan []byte) <-chan string {
	out := make(chan string)

	go func() {
		defer close(out)
		urls := parse(<-page)
		for _, u := range urls {
			select {
			case <-done:
				return
			case out <- u:
			}
		}
	}()
	return out
}

func isUrl(textUrl string) bool {
	_, err := url.ParseRequestURI(textUrl)
	if err == nil && strings.HasPrefix(textUrl, "http") {
		return true
	}
	return false
}
