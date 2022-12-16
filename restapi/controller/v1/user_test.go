package v1

import (
	"GinRESTful/restapi/test"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
)

type regLoginResponse struct {
	Code int          `json:"code"`
	Msg  string       `json:"msg"`
	Data regLoginData `json:"data"`
}

type regLoginData struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Token string `json:"token"`
}

func TestConRegister(t *testing.T) {
	// 初始化基础的测试环境
	test.InitTestBase()
	cases := []struct {
		UserName  string
		Password  string
		Password2 string
		CaptchaId string
		Captcha   string
		Code      uint
		Msg       string
		Name      string
	}{
		{UserName: "user10", Password: "admin12345", Password2: "admin12345", CaptchaId: "7rT6hsxxsFMmmMev3Qhh", Captcha: "92372", Code: 200, Msg: "注册成功", Name: "user10"},
		{UserName: "user10", Password: "admin12345", Password2: "admin12345", CaptchaId: "7rT6hsxxsFMmmMev3Qhh", Captcha: "92372", Code: 10017, Msg: "用户名已存在"},
		{UserName: "user11", Password: "admin12345", Password2: "admin12345678910", CaptchaId: "7rT6hsxxsFMmmMev3Qhh", Captcha: "92372", Code: 10016, Msg: "密码不一致"},
		{UserName: "user12", Password: "admin12345", Password2: "admin12345", CaptchaId: "7rT6hsxxsFMmmMev3Qhh", Captcha: "92372", Code: 10018, Msg: "注册失败"},
	}
	url := "/api/v1/user/register"
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST(url, ConRegister)
	for _, testCases := range cases {
		// 构造body Map
		bodyMap := make(map[string]string)
		bodyMap["user_name"] = testCases.UserName
		bodyMap["password"] = testCases.Password
		bodyMap["password2"] = testCases.Password2
		bodyMap["captcha_id"] = testCases.CaptchaId
		bodyMap["captcha"] = testCases.Captcha
		// 将bodyMap转化为json比特流
		jsonByte, _ := json.Marshal(bodyMap)
		// mock一个HTTP请求
		req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonByte))
		req.Header.Set("Content-Type", "application/json")
		// mock一个响应记录器
		w := httptest.NewRecorder()
		// 让server端处理mock请求并记录返回的响应内容
		r.ServeHTTP(w, req)
		// 校验系统状态码是否符合预期，系统状态码全部为200
		assert.Equal(t, 200, w.Code)
		// 解析并检验响应内容是否复合预期
		var resp regLoginResponse
		err := json.Unmarshal([]byte(w.Body.String()), &resp)
		assert.Nil(t, err)
		// 注册失败
		if testCases.Code != 200 {
			if testCases.Code == 10017 && resp.Code == 10017 {
				// 用户名已存在
				assert.Equal(t, testCases.Msg, resp.Msg)
			} else if testCases.Code == 10016 && resp.Code == 10016 {
				// 密码不一致
				assert.Equal(t, testCases.Msg, resp.Msg)
			} else if testCases.Code == 10018 && resp.Code == 10018 {
				// 注册失败
				assert.Equal(t, testCases.Msg, resp.Msg)
			}
		} else {
			// 校验是否正常返回注册的用户名
			assert.Equal(t, testCases.Name, resp.Data.Name)
			// 校验Token是否正确
			assert.Regexp(t, regexp.MustCompile("^[A-Za-z0-9-_=]+\\.[A-Za-z0-9-_=]+\\.?[A-Za-z0-9-_.+/=]*$"), resp.Data.Token)
		}
	}
}

func TestConLogin(t *testing.T) {
	// TODO
}

func TestConLogout(t *testing.T) {
	// TODO
}

func TestConGetMyselfInfo(t *testing.T) {
	// TODO
}

func TestConGetUserList(t *testing.T) {
	// TODO
}

func TestConModifyUserInfo(t *testing.T) {
	// TODO
}

func TestConDelUser(t *testing.T) {
	// TODO
}
