package model

import "spider/engine"

type SearchResult struct {
	Hits int
	Start int
	Items [] engine.Item
}