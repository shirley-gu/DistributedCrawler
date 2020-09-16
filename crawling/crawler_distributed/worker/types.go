package worker

import (
	"fmt"
	"log"
	"errors"
	"go_code/crawling/crawler/engine"
	"go_code/crawling/crawler_distributed/config"
)

type SerializedParser struct {
	Name string
	Args interface{}
}

//因为engine,Request里面有Parser是接口类型不能传输，所以要定义一个能够传输的Request
type Request struct {
	Url string
	Parser SerializedParser
}
//这个结构可以在网上传，要和engine.Request的结构相互转换
type ParseResult struct{
	Items []engine.Item
	Requests []Request
}

func SerializeRequest(r engine.Request) Request {
	name, args := r.Parser.Serialize()Request{
		Url: r.Url,
		Parser: SerializedParser{
			Name: name,
			Args: args,
		},
	}
}

func SerializeResult(r engine.ParseResult) ParseResult{
	result := ParseResult{
		Items: r.Items,
	}
	for _, req := range r.Requests {
		result.Requests = append(result.Requests, SerializeRequest(req))
	}
	return result
}

func DeserializeRequest(r Request) (engine.Request, error) {
	parser, err := deserializeParser(r.Parser)
	if err != nil {
		return engine.Request{}, err
	}
	return engine.Request{		
		Url: r.Url,
		Parser: parser, //r.Parser只是数据，需要能工作的parser，即下面的deserializeParser
	}, nil
}

func DeserializeResult(r ParseResult) engine.ParseResult {
	result := engine.ParseResult{
		Items: r.Items,
	}
	for _, req := range r.Requests {
		engineReq, err := DerializeRequest(req)
		if err != nil {
			log.Printf("error deserializing " + "request: %v", err)
			continue
		}
		result.Requests = append(result.Requests, engineReq)
	}
	return result
}

// type CrawlService struct {}
// func (CrawlService) Process(req engine.Request, result *engine.ParseResult) error {

// }

//两种方法：1.全局维护 2.switch
func deserializeParser(p SerializedParser) (engine.Parser,error) {
	switch p.Name{
		case config.ParseCityList:
			return engine.NewFuncParser(parser.ParseCityList,config.ParseCityList), nil
		case config.ParseCity:
			return engine.NewFuncParser(parser.ParseCity,config.ParseCity), nil
		case config.NilParser:
			return engine.NilParser{}, nil
		case config.ParseProfile:
			if userName, ok := p.Args.(string); ok {
				return parser.NewProfileParser(userName), nil
			} else {
				return nil, fmt.Errorf("invalid" + "arg: %v", p.Args)
			}
		default:	//这些人名都不认识
			return nil, errors.New("unknown parser name")
	}
}
