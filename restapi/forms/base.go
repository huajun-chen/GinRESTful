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

// CPUInfo CPU信息
type CPUInfo struct {
	CpuCounts      string `json:"cpu_counts"`       // CPU物理核心数
	CpuUsedpercent string `json:"cpu_used_percent"` // CPU使用率
}

// MemoryInfo 内存信息
type MemoryInfo struct {
	MemTotal       string `json:"mem_total"`        // 全部内存，单位GB
	MemUsed        string `json:"mem_used"`         // 已使用内存，单位GB
	MemFree        string `json:"mem_free"`         // 空闲内存，单位GB
	MemUsedPercent string `json:"mem_used_percent"` // 内存使用率
}

// DiskInfo 磁盘信息
type DiskInfo struct {
	DiskTotal string `json:"disk_total"` // 全部硬盘容量，单位GB
	DiskUsed  string `json:"disk_used"`  // 已使用硬盘，单位GB
	DiskFree  string `json:"disk_free"`  // 空闲硬盘，单位GB
}

// System 系统信息
type System struct {
	CPU    CPUInfo    `json:"cpu"`    // CPU
	Memory MemoryInfo `json:"memory"` // 内存
	Disk   DiskInfo   `json:"disk"`   // 磁盘
}
