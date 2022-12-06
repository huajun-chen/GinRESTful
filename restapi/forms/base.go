package forms

// PageForm 页数，每页数量
type PageForm struct {
	Page     int `form:"page" binding:"omitempty,gte=1,lte=10000"`      // 页数，第几页
	PageSize int `form:"page_size" binding:"omitempty,gte=1,lte=10000"` // 每页的数量
}

// IdForm ID
type IdForm struct {
	ID uint `uri:"id" binding:"required,gte=1,lte=100000000"` // 主键ID
}
