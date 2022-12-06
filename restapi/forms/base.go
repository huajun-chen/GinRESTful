package forms

// PageForm 页数，每页数量
type PageForm struct {
	Page     int `form:"page"`      // 页数，第几页
	PageSize int `form:"page_size"` // 每页的数量
}

// IdForm ID
type IdForm struct {
	ID uint `uri:"id"`
}
