package main

import (
	"fmt"
	"flag"
	"go_code/crawling/crawler/engine"
	"go_code/crawling/crawler/scheduler"
	"go_code/crawling/crawler/zhenai/parser"
	"go_code/crawling/crawler_distributed/config"
	itemsaver "go_code/crawling/crawler_distributed/persist"
	"go_code/crawling/crawler_distributed/rpcsupport"
	worker "go_code/crawling/crawler_distributed/worker"
)

var (
	itemSaverHost = flag.String("itemsaver_host", "", "itemsaver host")
	workerHosts = flag.String("worker_hosts", "", "worker hosts (comma separated)")
)

func main(){
	flag.Parse()
	itemChan, err := itemsaver.ItemSaver(fmt.Sprintf(":%d", *itemSaverHost))
	//itemChan, err := itemsaver.ItemSaver(fmt.Sprintf(":%d", config.ItemSaverPort))
	if err!= nil{
		panic(err)
	}

	pool := createClientPool(strings.Split(*workerHosts, ","))

	processor:=worker.CreateProcessor(pool)
	
	e:=engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan: itemChan,
		RequestProcessorCreator: processor,
	}

	e.Run(engine.Request){
		Url:"http://www.zhenai.com/zhenhun"
		Parser:engine.NewFuncParser(

		)
	}
}

//通过createclientpool 形成pool 源源不断的通过channel发送给client
func createClientPool(hosts []string) chan *rpc.Client{
	var clients []*rpc.Client	//建立clients，把clients全部连接起来，clients是createclientPool私有的数据，不共享
	for _,h := range hosts{
		clients,err := rpcsuppport.NewClient(h)
		if err == nil{
			clients = append(clients, client)
			log.Printf("Connected to %s", h)
		}else{
			log.Printf("err connection to %s: %v", h, err)
		}
	}

	//分发在gorutine里面
	//通过channel(消息传递)来共享数据，clients通过消息传递共享给外面的worker，worker就可以去连接了
	out:=make(chan *rpc.Client)
	go func(){
		for{
			for _, client := range clients{
			out <- client	//
		}
	}()
	return out
}