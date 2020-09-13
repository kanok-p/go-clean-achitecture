package request

import "github.com/kanok-p/go-clean-architecture/util/pagination"

type GetListInput struct {
	*pagination.Pagination
	Limit  int64
	Page   int64
	Offset int64
	Search string
}
