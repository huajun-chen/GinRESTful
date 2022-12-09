package utils

import (
	"GinRESTful/restapi/forms"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"strconv"
	"time"
)

// CPUInfo 系统CPU信息
// 参数：
//		无
// 返回值：
//		forms.CPU：CPU核心数，使用率
//		error：错误信息
func CPUInfo() (forms.CPUReturn, error) {
	cpuStruct := forms.CPUReturn{}
	// CPU核心数，参数true：逻辑内核，参数false：物理内核
	numCPUs, err := cpu.Counts(false)
	if err != nil {
		return cpuStruct, err
	}
	// 使用率保留2位小数，获取1s内的CPU使用率信息，太短不准确，也可以获几秒内的，但这样延迟太大
	cpuUsage, err := cpu.Percent(time.Second, false)
	if err != nil {
		return cpuStruct, err
	}

	cpuStruct.CpuCounts = strconv.Itoa(numCPUs)
	cpuStruct.CpuUsedpercent = strconv.FormatFloat(cpuUsage[0], 'f', 2, 64)
	// 可以通过fmt.Sprintf("%.2f", num)将float64转为字符串并保留2位小数
	// 可以通过strconv.FormatFloat(num,'f',2,64)将float64转为字符串并保留2位小数
	// 经过多次自测之后strconv.FormatFloat(num,'f',2,64)性能更好一些

	return cpuStruct, nil
}

// MemInfo 系统内存信息
// 参数：
//		无
// 返回值：
//		forms.Memory：内存全部，已使用，未使用，使用率
//		error：错误信息
func MemInfo() (forms.MemoryReturn, error) {
	memStruct := forms.MemoryReturn{}
	// 获取内存信息
	memUsage, err := mem.SwapMemory()
	if err != nil {
		return memStruct, err
	}
	// 全部的，将字节数转换为GB
	memTotal := float64(memUsage.Total) / 1024 / 1024 / 1024
	// 已使用，将字节数转换为GB
	memUsed := float64(memUsage.Used) / 1024 / 1024 / 1024
	// 未使用，将字节数转换为GB
	memFree := float64(memUsage.Free) / 1024 / 1024 / 1024
	// 使用率，保留2位小数
	memUsedpercent := memUsage.UsedPercent

	memStruct.MemTotal = strconv.FormatFloat(memTotal, 'f', 2, 64)
	memStruct.MemUsed = strconv.FormatFloat(memUsed, 'f', 2, 64)
	memStruct.MemFree = strconv.FormatFloat(memFree, 'f', 2, 64)
	memStruct.MemUsedPercent = strconv.FormatFloat(memUsedpercent, 'f', 2, 64)

	return memStruct, nil
}

// DiskInfo 系统硬盘信息
// 参数：
//		无
// 返回值：
//		forms.Disk：内存全部容量，已使用容量，未使用容量
//		error：错误信息
func DiskInfo() (forms.DiskReturn, error) {
	diskStruct := forms.DiskReturn{}
	// 获取磁盘信息
	diskUsage, err := disk.Usage("/")
	if err != nil {
		return diskStruct, err
	}
	// 全部的，将字节数转换为GB
	diskTotal := float64(diskUsage.Total) / 1024 / 1024 / 1024
	// 已使用，将字节数转换为GB
	diskUsed := float64(diskUsage.Used) / 1024 / 1024 / 1024
	// 未使用，将字节数转换为GB
	diskFree := float64(diskUsage.Free) / 1024 / 1024 / 1024

	diskStruct.DiskTotal = strconv.FormatFloat(diskTotal, 'f', 2, 64)
	diskStruct.DiskUsed = strconv.FormatFloat(diskUsed, 'f', 2, 64)
	diskStruct.DiskFree = strconv.FormatFloat(diskFree, 'f', 2, 64)

	return diskStruct, nil
}
