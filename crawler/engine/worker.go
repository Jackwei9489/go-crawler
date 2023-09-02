package engine

import (
	"go-crawler/crawler/fetcher"
	"log"
)

func worker(r Request) (ParseResult, error) {
	//log.Printf("Fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("error fetching, ur=%s, err=%v", r.Url, err)
		return NilParser(body), err
	}

	return r.ParseFunc(body, r.Url), nil
}
