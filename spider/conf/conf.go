package conf

type EsConfType map[string]string

var EsConf = EsConfType{
	"Url":   "http://39.100.78.46:9200",
	"Index": "db1",
	"Type":  "user",
}
