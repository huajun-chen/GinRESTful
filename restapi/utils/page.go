package utils

import "GinRESTful/restapi/global"

// PageZero 判断page和page_size参数是否为0
func PageZero(page, pageSize int) (int, int) {
	if page == 0 || pageSize == 0 {
		page = global.Page
		pageSize = global.PageSize
	}
	return page, pageSize
}

// LimitResult 计算limit
func LimitResult(pageSize int) int {
	return pageSize
}

// OffsetResult 计算offset
func OffsetResult(page, pageSize int) int {
	limit := LimitResult(pageSize)
	return (page - 1) * limit
}
