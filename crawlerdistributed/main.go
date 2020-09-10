
func main(){
	itemChan,err:=itemsaver.ItemSaver(
		fmt.Sprintf("%d",config.ItemSaverPort))
		if err!= nil{
			panic(err)
		}
	pool := createClientPool()
	processor:=worker.CreateProcessor()
	
	e:=engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkerCount:100,
		ItemChan:itemChan,
		RequestProcessorCreator:processor,
	}

	e.Run(engine.Request){
		Url:"http://www.zhenai.com/zhenhun"
		Parser:engine.NewFuncParser(

		)
	}
}
//通过createclientpool 形成pool 源源不断的通过channel发送给client
func createClientPool(hosts []string) chan *rpc.Client{
	var clients []*rpc.Client
	for _,h:=range hosts{
		clients,err:= rpcsuppport.NewClient(h)
		if err == nil{
			clients= append(clients,client)
			log.Printf("Connected to %s",h)
		}else{
			log.Printf("err connection to %s:%v",h,err)
		}
	}
	out:=make(chan *rpc.Client)
	go func(){
		for{
			for _,client:=range clients{
			out <- client
		}
	}
	}()
	return out
}