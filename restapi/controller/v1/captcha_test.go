package v1

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type captchaResponse struct {
	Code int         `json:"code"` // 自定义状态码
	Msg  string      `json:"msg"`  // 信息
	Data captchaData `json:"data"` // 数据
}

type captchaData struct {
	CaptchaID   string `json:"captcha_id"`   // 验证码ID
	CaptchaPath string `json:"captcha_path"` // 验证码bs64码
}

func TestGetCaptcha(t *testing.T) {
	url := "/api/v1/base/captcha"
	// 设置Gin框架运行模式（ReleaseMode：生产环境、TestMode：测试环境）
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET(url, ConGetCaptcha)
	cases := []captchaResponse{
		{Code: 10007, Msg: "生成验证码失败"},
		{Code: 200, Data: captchaData{CaptchaID: "", CaptchaPath: ""}},
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
		var resp captchaResponse
		err := json.Unmarshal([]byte(w.Body.String()), &resp)
		assert.Nil(t, err)
		// 生成验证码失败
		if testCases.Code != 200 {
			if testCases.Code == 10007 && resp.Code == 10007 {
				assert.Equal(t, testCases.Msg, resp.Msg)
			}
		} else {
			// 验证自定义状态码是否为200
			assert.Equal(t, testCases.Code, resp.Code)
			// 验证码ID的长度是否为20
			assert.Equal(t, 20, len(resp.Data.CaptchaID))
			// 验证码图片的base64码是否包含：data:image/png;base64
			assert.Contains(t, resp.Data.CaptchaPath, "data:image/png;base64", "strings contains substr")
		}
	}
}
