package parser

import (
	"go_code/crawling/crawler/types"
	"regexp"
)

//var CityListRe = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)

// 获取城市列表
const cityListRe = `< a href=" "[^>]*(>[^<]+)</ a>`

func ParseCityList (contents []byte) engine.ParseResult{
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)
	//limit:= 10
	result := engine.ParseResult{}
	for _,m := range matches{
		//result.Items = append(result.Items,"City "+string(m[2]))
		result.Requests = append(result.Requests,engine.Request{
			Url : string(m[1]),
			Parser: engine.CreateFuncParser(ParseCity,"ParseCity") 
		})
	}
	return result
}

// func ParseCityList(contents []byte) types.ParseResult {
// 	matches := CityListRe.FindAllSubmatch(contents, -1)
// 	result := types.ParseResult{}
// 	for _, m := range matches {
// 		result.Requests = append(
// 			result.Requests,
// 			types.Request{Url: string(m[1]), ParseFunc: ParseCityUserList})
// 	}
// 	return result
// }