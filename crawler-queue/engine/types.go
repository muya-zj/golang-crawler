package engine

type ParseFunc func(contents []byte, url string) ParseResult

type Parser interface {
	Parse(contents []byte, url string) ParseResult
	Serialize() (name string, args interface{})
}

//解析请求struct
type Request struct {
	Url    string //要解析的url
	Parser Parser //该url对应的解析函数
}

type SerializedParser struct {
	Name string      //函数名
	Args interface{} //参数
}

//{"ParserCityList",nil},{"ProfileParser",userName}

//解析后返回的结果集struct
type ParseResult struct {
	Requests []Request //要请求的内容
	Items    []Item    //结果集的具体内容
}

type Item struct {
	Url     string      //人物的url
	Id      string      //人物的ID,去重时用的,也用作ElasticSearch的Id
	Type    string      //ElasticSearch的table name
	Payload interface{} //具体数据
}

type NilParser struct{}

func (NilParser) Parse(
	_ []byte, _ string) ParseResult {
	return ParseResult{}
}

func (NilParser) Serialize() (
	name string, args interface{}) {
	return "NilParser", nil
}

type FuncParser struct {
	parser ParseFunc
	name   string
}

func (f *FuncParser) Parse(contents []byte, url string) ParseResult {
	return f.parser(contents, url)
}

func (f *FuncParser) Serialize() (
	name string, args interface{}) {
	return f.name, nil
}

func NewFuncParser(p ParseFunc, name string) *FuncParser {
	return &FuncParser{
		parser: p,
		name:   name,
	}
}
