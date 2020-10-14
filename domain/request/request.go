package request


type PageOption struct {
	Page    int64  `form:"page" binding:"min=0"`
	PerPage int64  `form:"per_page" binding:"min=0"`
	Search  string `form:"search"`
}
