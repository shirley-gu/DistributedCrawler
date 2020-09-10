package main
import(
	"goproject/crawlerdistributed/rpc"
	"goproject/crawlerdistributed/persist"
)

func main(){
	log.Fatal(serveRpc (":1234","dating_profile"))	
}

func serveRpc(host ,index string) error{
	client,err:= elastic.NewClient(
		elastic.SetSniff(false))
	if err!=nil{
		return err
	}
	return rpcsupport.ServeRpc( host ,persist.ItemSaverService{
		Client:client,
		Index: index,
	})
	return nil
}