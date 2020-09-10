package rpcsupport
import(
	"net/rpc"
	"net"
	"net/rpc/jsonrpc"
	"goproject/crawler/rpc"
)
func ServerRpc(host stirng, service interface{}) error{
	rpc.Register(rpcdemo.DemoService{})

	listener,err := net.Listen("tcp",host)
	if err!= nil{
		panic(err)
	}
	for{
		conn,err:= listen.Accept()
		if err!= nil{
			log.Printf("accept error:%v",err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
	return nil
}

func NewClient(host string)(*rpc.Client,error){
	conn,err:= net.Dial("tcp",":1234")
	if err!= nil{
		//panic(err)
		return nil,err
	}
	return jsonrpc.NewClient(conn),nil
}