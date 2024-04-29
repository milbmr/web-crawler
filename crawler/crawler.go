package crawler

import (
	"bytes"
	"fmt"
	"log"
	"strings"

	"golang.org/x/net/html"
)

func Parser(page []byte) {
	doc, err := html.Parse(bytes.NewReader(page))
	if err != nil {
		log.Fatal(err)
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					if strings.HasPrefix(a.Val, "http") {
						fmt.Println(a.Val)
					}
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
