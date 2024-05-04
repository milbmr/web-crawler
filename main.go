package main

import (
	"io"
	"log"
	"net/http"

	"github.com/milbmr/web-crawler/crawler"
)

func main() {
  url := make(chan string)
  done := make(chan interface{})
  defer close(done)

  crawler.Crawl()
}
