package search

import (
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/kanok-p/go-clean-architecture/util/search"
)

func parseMongoOrder(sortStr string) int64 {
	if strings.ToUpper(sortStr) == "DESC" {
		return -1
	}

	return 1

}

func ToMongoSort(sorts []*search.KeyValue) bson.D {
	sortQuery := make(bson.D, 0, len(sorts))
	for _, sort := range sorts {
		sortQuery = append(sortQuery, bson.E{
			Key:   sort.Key,
			Value: parseMongoOrder(sort.Value),
		})
	}

	return sortQuery
}

func ToMongoFilter(filters []*search.KeyValue) bson.D {
	filterQuery := make(bson.D, 0, len(filters))
	for _, filter := range filters {
		filterQuery = append(filterQuery, bson.E{
			Key: filter.Key,
			Value: bson.M{"$regex": primitive.Regex{
				Pattern: filter.Value,
				Options: "i",
			}},
		})
	}

	return filterQuery
}
