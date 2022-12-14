package v1

import (
	"GinRESTful/restapi/global"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type testResponse struct {
	Code int    `json:"code"` // 自定义状态码
	Msg  string `json:"msg"`  // 信息
	Data data   `json:"data"` // 数据
}

type data struct {
	CaptchaID   string `json:"captcha_id"`   // 验证码ID
	CaptchaPath string `json:"captcha_path"` // 验证码bs64码
}

func TestGetCaptcha(t *testing.T) {
	url := "/api/v1/base/captcha"
	r := gin.Default()
	r.GET(url, ConGetCaptcha)
	cases := []testResponse{
		{Code: 10007, Msg: "生成验证码失败"},
		{Code: 10007, Msg: "Failed to Generate Verification Code"},
		{Code: 200, Data: data{CaptchaID: "", CaptchaPath: ""}},
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
		var resp testResponse
		err := json.Unmarshal([]byte(w.Body.String()), &resp)
		assert.Nil(t, err)
		// 生成验证码失败
		if testCases.Code == 10007 {
			if resp.Code == 1007 {
				// Msg信息是否为：生成验证码失败
				settLang := global.Settings.Language.LanguageType
				if settLang == "zh-CN" {
					// 中文
					assert.Equal(t, testCases.Msg, resp.Msg)
				} else if settLang == "en-US" {
					// 英文
					assert.Equal(t, testCases.Msg, resp.Msg)
				}
			}
		} else {
			// 生成验证码正确
			if resp.Code == 200 {
				// 验证码ID的长度是否为20
				assert.Equal(t, 20, len(resp.Data.CaptchaID))
				// 验证码图片的base64码是否包含：data:image/png;base64
				assert.Contains(t, resp.Data.CaptchaPath, "data:image/png;base64", "strings contains substr")
			}
		}
	}
}
