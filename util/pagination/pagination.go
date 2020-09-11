package pagination

type Pagination struct {
	Page  int64
	Limit int64
}

type Output struct {
	Total   int64
	Limit   int64
	Page    int64
	HasNext bool
}

func (p *Pagination) GetPage() int64 {
	if p == nil {
		return 0
	}

	return p.Page
}

func (p *Pagination) GetOffset() int64 {
	var offset int64
	if p == nil {
		return 0
	}

	if p.Page != 0 {
		offset = (p.Page - 1) * p.Limit
	}

	return offset
}

func (p *Pagination) GetLimit() int64 {
	if p == nil {
		return 0
	}

	return p.Limit
}

func GetPaginationSkipAndHasNext(page, limit, total int64) (int64, bool) {
	hasNext := true

	skip := (page - 1) * limit

	if (page * limit) >= total {
		hasNext = false
	}

	return skip, hasNext
}
