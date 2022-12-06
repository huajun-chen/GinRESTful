package utils

import "GinRESTful/restapi/global"

// PageZero 判断page和page_size参数是否为0
// 参数：
//		page：页数/第几页
//		pageSize：每页的数量
// 返回值：
//		int：处理后的页数
//		int：处理后的每页的数量
func PageZero(page, pageSize int) (int, int) {
	if page == 0 || pageSize == 0 {
		page = global.Settings.Page
		pageSize = global.Settings.PageSize
	}
	return page, pageSize
}

// LimitResult 计算limit
// 参数：
//		pageSize：每页的数量
// 返回值：
//		int：数据库分页查询需要的limit值
func LimitResult(pageSize int) int {
	return pageSize
}

// OffsetResult 计算offset
// 参数：
//		page：页数/第几页
//		pageSize：每页的数量
// 返回值：
//		int：数据库分页查询需要的offset值
func OffsetResult(page, pageSize int) int {
	limit := LimitResult(pageSize)
	return (page - 1) * limit
}
