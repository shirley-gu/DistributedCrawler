package types

//type ParseFunc func(contents []byte, url string) ParseResult
//接收contents和url产生ParseResult
type Parser interface{
	Parse(contents []byte, url string) ParseResult
	Serialize() (name string, args interface{})
}

type Request struct {
	Url       string
	Parser    Parser
}

// type SerializedParser struct {
// 	Name string //函数的名字
// 	Args interface{}  //参数
// }

type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

// func NilParser([]byte) ParseResult {
// 	return ParseResult{}
// }

type Item struct {
	Url string
	Type string
	Id string
	Payload interface{}
}

type NilParser struct{}

func (NilParser) Parse() (contents []byte, url string) ParseResult{
	return ParseResult{}
}

func (NilParser) Serialize() (name string, args interafce{}) {
	return "Nil Parser", nil
}

type FuncParser struct {
	parser ParseFunc
	Name String
}

func (f *FuncParser) Parse(contents []byte, url string) ParseResult{
	return f.parser(contents, url)
}

func (f *FuncParser) Serialize() (name string, args interface{}) ParseResult{
	return f.name, nil
}

func NewFuncParser(p ParserFunc, name string) FuncParser {
	return FuncParser{
		parser: p,
		name : name,
	}
}

type ProfileParser struct {
	userName string
}

func newProfileParser(name string) *ProfileParser{
	return &ProfileParser{
		userName:name,
	}
}

func (p *ProfileParser) Parse(contents []byte, url string) engine.ParseResult{
	return parseProfile(contents, url, p.userName)
}

func (p *ProfileParser) Serialize() (name string, args interface{}){
	return "ProfileParser", p.userName
}