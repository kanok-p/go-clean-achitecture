package search

import (
	"log"
	"strings"

	"github.com/kanok-p/go-clean-architecture/util/pagination"
)

type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Input struct {
	*pagination.Pagination
	FilterValues []*KeyValue
	SortValues   []*KeyValue
}

type Option interface {
	GetPage() int64
	GetLimit() int64
	GetSort() string
	GetFilter() []*KeyValue
}

type Search struct {
	*pagination.Pagination
	Sort string `json:"sort" form:"sort,default=createdAt:desc"`
}

func (s *Search) GetSort() string {
	return s.Sort
}

func New(so Option) *Input {
	sorts := ParseSortFromRawQuery(so.GetSort())
	return &Input{
		Pagination: &pagination.Pagination{
			Page:  so.GetPage(),
			Limit: so.GetLimit(),
		},
		FilterValues: so.GetFilter(),
		SortValues:   sorts,
	}
}

func ParseSortFromRawQuery(rawQuery string) []*KeyValue {
	var sorts []*KeyValue
	if rawQuery == "" {
		return sorts
	}

	sortArr := strings.Split(rawQuery, ",")
	log.Println("SortArr : ", sortArr)
	for _, s := range sortArr {
		sortOption := strings.Split(s, ":")
		if len(sortOption) == 1 {
			sortOption = append(sortOption, "asc")
		}

		sorts = append(sorts, &KeyValue{
			Key:   sortOption[0],
			Value: sortOption[1],
		})
	}

	return sorts
}
