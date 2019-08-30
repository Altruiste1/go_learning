package engine

import (
	"github.com/Altruiste1/go_learning/crawler/Concurrent_crawler/fetcher"
	"github.com/labstack/gommon/log"
)

type SimpleEngine struct {}
func (s *SimpleEngine)Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		parseResult,err:=worker(r)
		if err!=nil{
			continue
		}

		requests = append(requests,
			parseResult.Requests...)
		for _, item := range parseResult.Items {
			log.Printf("Got item %s", item)
		}
	}
}

func worker(r Request)(ParseResult,error){
	log.Printf("Fetching %s\n", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher:error "+
			"fetching url %s,%v", r.Url, err)
		return ParseResult{},err
	}

	return r.ParserFunc(body),nil
}
