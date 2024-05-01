package crawler

import (
	"bytes"
	"fmt"
	"log"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func Parser(page []byte) string {
	doc, err := html.Parse(bytes.NewReader(page))
	if err != nil {
		log.Fatal(err)
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" && isUrl(a.Val) {
					fmt.Println(a.Val)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
}

func isUrl(textUrl string) bool {
	_, err := url.ParseRequestURI(textUrl)
	if err == nil && strings.HasPrefix(textUrl, "http") {
		return true
	}
	return false
}
