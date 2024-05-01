package main

import (
	"io"
	"log"
	"net/http"

	"github.com/milbmr/web-crawler/crawler"
)

func main() {
	res, err := http.Get("https://yts.mx/")
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err) }
	res.Body.Close()
  crawler.Parser(body)
}
