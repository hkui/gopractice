package model

import (
	"spider/model"
)

type SearchResult struct {
	Hits int
	Start int
	Items [] model.Profile
}