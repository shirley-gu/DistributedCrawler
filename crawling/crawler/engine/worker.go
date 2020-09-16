package engine

import (
	"log"
	"go_code/crawling/crawler/fetcher"
)

//输入+输出+error
//需要把这个worker包装
func worker(r Request) (ParseResult, error) {
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error" + "fetching url %s: %v", r.Url, err)
		return ParseResult{}, err
	}
	//return r.ParserFunc(body, r.Url), nil
	return r.Parser.Parse(body, r.Url), nil
}