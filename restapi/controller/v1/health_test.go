package v1

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
)

type healthResponse struct {
	Code int        `json:"code"` // 自定义状态码
	Msg  string     `json:"msg"`  // 信息
	Data healthData `json:"data"` // 数据
}

type cpu struct {
	CPUCounts      string `json:"cpu_counts"`
	CPUUsedPercent string `json:"cpu_used_percent"`
}

type memory struct {
	MemTotal       string `json:"mem_total"`
	MemUsed        string `json:"mem_used"`
	MemFree        string `json:"mem_free"`
	MemUsedPercent string `json:"mem_used_percent"`
}

type disk struct {
	DiskTotal string `json:"disk_total"`
	DiskUsed  string `json:"disk_used"`
	DiskFree  string `json:"disk_free"`
}

type healthData struct {
	CPU    cpu    `json:"cpu"`
	Memory memory `json:"memory"`
	Disk   disk   `json:"disk"`
}

func TestConGetSystemInfo(t *testing.T) {
	url := "/api/v1/base/health"
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET(url, ConGetSystemInfo)
	cases := []healthResponse{
		{Code: 10023, Msg: "获取CPU信息失败"},
		{Code: 10024, Msg: "获取内存信息失败"},
		{Code: 10025, Msg: "获取硬盘信息失败"},
		{Code: 200},
	}
	for _, testCases := range cases {
		// mock一个HTTP请求
		req := httptest.NewRequest(http.MethodGet, url, nil)
		// mock一个响应记录器
		w := httptest.NewRecorder()
		// 让server端处理mock请求并记录返回的响应内容
		r.ServeHTTP(w, req)
		// 校验系统状态码是否符合预期，系统状态码全部为200
		assert.Equal(t, 200, w.Code)

		// 解析并检验响应内容是否复合预期
		var resp healthResponse
		err := json.Unmarshal([]byte(w.Body.String()), &resp)
		assert.Nil(t, err)
		// 获取信息失败
		if testCases.Code != 200 {
			if testCases.Code == 10023 && resp.Code == 10023 {
				assert.Equal(t, testCases.Msg, resp.Msg)
			} else if testCases.Code == 10024 && resp.Code == 10024 {
				assert.Equal(t, testCases.Msg, resp.Msg)
			} else if testCases.Code == 10025 && resp.Code == 10025 {
				assert.Equal(t, testCases.Msg, resp.Msg)
			}
		} else {
			// 获取信息正确
			assert.Equal(t, testCases.Code, resp.Code)
			// 正则匹配cpu核心数cpu_counts必须为数字，并且数字大于0
			assert.Regexp(t, regexp.MustCompile("^[1-9]\\d*$"), resp.Data.CPU.CPUCounts)
			// 正则匹配内存容量mem_total
			assert.Regexp(t, regexp.MustCompile("^[1-9]\\d*(\\.\\d{1,2})?$"), resp.Data.Memory.MemTotal)
			// 正则匹配硬盘容量disk_total
			assert.Regexp(t, regexp.MustCompile("^[1-9]\\d*(\\.\\d{1,2})?$"), resp.Data.Disk.DiskTotal)
		}
	}
}
