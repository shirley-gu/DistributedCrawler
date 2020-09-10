package worker
import(
	"fmt"
	"log"
)
var port = flag.Int("port",0,"the port for me me to listein on")
func main(){
	flag.Parse()
	if *port == 0{
		fmt.Println("must specify s port")
		return
	}
	log.Fatal(rpcsupport.ServeRpc(
		fmt.Sprintf("%d",*port),
		worker.CrawlService{}))
}