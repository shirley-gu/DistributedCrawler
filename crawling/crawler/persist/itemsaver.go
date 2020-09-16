package persist
import (
	"log"
)
func ItemSaver() chan interface{
	out := make(chan interface{})
	go func(){
		itemCount := 0
		for{
			item := <- out
			log.Printf("ItemSaver:got item"+"%d:%v",itemCount,item)
			itemCount++
		}
	}()
	return out
}

func Save(client *elastic.Client, index string,item engine.Item) error{
	if item.Type == ""{
		return errors.New("must supply Type")
	}
	indexService := client.Index().Index(index).Type(item.Type).BodyJson(item)
	if item.Id != ""{
		indexService.Id(item.Id)
	}
	_, err:= indexService.
}