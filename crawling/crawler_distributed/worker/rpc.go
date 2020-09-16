package worker

import (
	"go_code/crawling/crawler/engine"
)

type CrawlService struct {}

//func (CrawlService) Process(req engine.Request, result *engine.ParseResult) error {
func (CrawlService) Process(req Request, result *ParseResult) error {
	engineReq, err := DeserializeRequest(req)
	if err != nil {
		return err
	}
	engineResult, err := engine.Worker(engineReq)
	if err != nil {
		return err
	}
	*result = SerializeResult(engineResult)
	return nil
}
