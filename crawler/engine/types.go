package engine

type Request struct {
	Url       string
	ParseFunc ParserFunc
}

type ParserFunc func([]byte, string) ParseResult

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Id      string      `json:"id"`
	Url     string      `json:"url"`
	Payload interface{} `json:"payload"`
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}
