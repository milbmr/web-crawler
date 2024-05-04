package crawler

import (
	"bytes"
	"fmt"
	"log"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func Parser(done <-chan interface{}, page <-chan []byte) <-chan string {
	urls := make(chan string)

	go func() {
    defer close(urls)
		doc, err := html.Parse(bytes.NewReader(<-page))
		if err != nil {
			log.Fatal(err)
		}
		var f func(*html.Node)
		f = func(n *html.Node) {
			if n.Type == html.ElementNode && n.Data == "a" {
				for _, a := range n.Attr {
					if a.Key == "href" && isUrl(a.Val) {
            select
						break
					}
				}
			}
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				f(c)
			}
		}
		f(doc)
	}()
  return urls
}

func isUrl(textUrl string) bool {
	_, err := url.ParseRequestURI(textUrl)
	if err == nil && strings.HasPrefix(textUrl, "http") {
		return true
	}
	return false
}
