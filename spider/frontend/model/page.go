package model

type SearchResult struct {
	Hits int64
	Start int
	Items []interface{}
	Query string
	PageHtml string //页码区域html代码
}