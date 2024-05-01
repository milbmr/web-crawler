package crawler

import (
	"io"
	"log"
	"net/http"
)

func Crawler(baseUrl string, depth int, found chan string) {
	if depth <= 0 {
		return
	}

	res, err := http.Get("https://yts.mx/")
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	res.Body.Close()
	Parser(body)
}
