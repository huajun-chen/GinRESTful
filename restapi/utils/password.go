package utils

import (
	"GinRESTful/restapi/global"
	"golang.org/x/crypto/bcrypt"
)

// SetPassword 设置密码，加密密码
// 参数：
//		password：未加密的原始密码
// 返回值：
//		string：加密后的密码
//		error：错误信息
func SetPassword(password string) (string, error) {
	userInfo := global.Settings.UserInfo
	pwdBytes, err := bcrypt.GenerateFromPassword([]byte(password), userInfo.PwdEncDiff)
	return string(pwdBytes), err
}

// CheckPassword 校验密码
// 参数：
//		memoryPwd：数据库已经存在的密码
//		enterPwd：当前输入的密码
// 返回值：
//		bool：密码是否一致
func CheckPassword(memoryPwd, enterPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(memoryPwd), []byte(enterPwd))
	return err == nil
}
