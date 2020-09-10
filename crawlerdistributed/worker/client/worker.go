package worker

import (
	"fmt"
	"net/rpc"
	"goproject/crawler/engine"
)
func CreateProcessor(clientChan chan *rpc.Client)engine.Processor{
	// client,err:=rpcsupport.NewClient(
	// 	fmt.Spritnf("%d",config.WorkerPort0))
	// 	if err!= nil{
	// 		return nil.err
	// 	}
		return func(
			req engine.Request)(
			engine.ParseResult,error){
				sReq:=worker.SerializeRequest(req)
					
				var sResult worker.ParseResult
				c := <-clientChan
				err := c.Call(config.CrawlServiceRpc,
				sReq,&sREsult)
				if err!= nil{
					return engine.ParseResult{},err
				}
				return worker.DeserializeResult(sResult),
				nil
				}
}
