package server

import(
	"go_code/crawlering/crawler_distributed/rpcsupport"
	"go_code/crawlering/crawler_distributed/persist"
)

var port = flag.Int("port", 0, "the port for me me to listein on")

func main() {
	flag.Parse()
	if *port == 0{
		fmt.Println("must specify a port")
		return
	}

	log.Fatal(serveRpc (":%d", *port), config.ElasticIndex)	
}

// func main(){
// 	log.Fatal(serveRpc (":1234","dating_profile"))	
// }

func serveRpc(host ,index string) error{
	client,err:= elastic.NewClient(
		elastic.SetSniff(false))
	if err!=nil{
		return err
	}
	return rpcsupport.ServeRpc( host , persist.ItemSaverService{
		Client: client,
		Index: index,
	})
	return nil
}