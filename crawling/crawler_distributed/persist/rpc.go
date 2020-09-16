package persist

import (
	"elastic"
	"go_code/crawling/crawler/engine"
	"go_code/crawling/crawler/persist"
)

type ItemSeverService struct{
	Client *elastic.Client
	Index string
}
func (s *ItemSeverService) Save(item engine.Itme,result *string) error{
	persist.Save(s.Client,s.Index, item)
	if err==nil{
		*result = "ok"
	}
	return err
}