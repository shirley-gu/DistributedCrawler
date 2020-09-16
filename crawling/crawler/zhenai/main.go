package main

import (
	"go_code/crawling/crawler/engine"
	"go_code/crawling/crawler_distributed/config"
	"go_code/crawling/crawler/scheduler"
	"go_code/crawling/crawler/types"
	"go_code/crawling/crawler/zhenai/parser"
)

func main() {
	itemChan,err:=persist.ItemSaver("dating_profile")
	if err!= nil{
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:itemChan,
		RequestProcessor:engine.Worker,
	}
	e.Run(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		Parser:engine.NewFuncParser(
		parser.ParseCityList,
		config.ParseCityList),
	})
}