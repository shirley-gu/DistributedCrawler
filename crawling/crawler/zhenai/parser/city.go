package parser

import (
 "fmt"
 "go_code/crawling/crawler/engine"
 "regexp"
 //"strconv"
)
 

const cityRe = `< a href=" "[^>]*(>[^<]+)</ a>`
func ParseCity(
	contents []byte) engine.ParseResult{
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _,m:= range matches{
		name := string(m[2])
		result.Items = append(result.Items,"User"+name)
		result.Requests = append(result.Requests,engine.Request{
			Url : string(m[1]),
			Parser : engine.NewFuncParser(){
				ParseCity,"ParseCity"),
			},
		})
		fmt.Printf("City: %s, URL: %s\n",m[2],m[1])
	}
	return result
}