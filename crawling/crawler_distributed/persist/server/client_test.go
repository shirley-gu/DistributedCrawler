package server

import (
	"testing"
	"time"
)
func serveRpc(
	host string, service interface{}) error{
		
	}
func TestItemSaver(t *testing.T){
	const host = ":1234"
	//start 
	go serveRpc(host,"test1")
	time.Sleep(time.Second)
	//start ItemSaverClient
	client,err := rpcsupport.NewClient(host)
	if err != nil{
		panic(err)
	}
	//Call save
	item := engine.Item{
		Url:"http://album.zhenai.com/u/108906739",
		Type:"zhenai",
		Id:"108906739",
		Payload:model.Profile{
			Name  string
			Gender  string
			Age  int
			Height  int
			Weight  int 
			Income  string 
			Marriage  string 
			Education  string 
			Occupation  string
			Hokou   string  
			Xinzuo  string 
			House  string 
			Car  string 	
		}
	}
	result := ""
	err=client.Call("ItemSaverService.Save",item,&result)
	if err!=nil || result != ok{
		t.Errorf("result:%s; err:%s",result,err)
	}
}