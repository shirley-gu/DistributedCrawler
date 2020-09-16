package worker
import (
	"fmt"
	"log"
	"go_code/crawling/crawler_distributed/config"
	"go_code/crawling/crawler_distributed/rpcsupport"
	"go_code/crawling/crawler_distributed/worker"
)

//建立命令行参数(整数类型)
var port = flag.Int("port", 0, "the port for me me to listein on")

func main(){
	flag.Parse()
	if *port == 0{
		fmt.Println("must specify a port")
		return
	}
	log.Fatal(rpcsupport.ServeRpc(
		fmt.Sprintf(":%d", *port),
		worker.CrawlService{}))
}

//测试：go run worker.go --port=9000