package utils

import (
	"GinRESTful/restapi/global"
	"golang.org/x/crypto/bcrypt"
)

// SetPassword 设置密码，加密密码
func SetPassword(password string) (string, error) {
	userInfo := global.Settings.UserInfo
	pwdBytes, err := bcrypt.GenerateFromPassword([]byte(password), userInfo.PwdEncDiff)
	return string(pwdBytes), err
}

// CheckPassword 校验密码
func CheckPassword(memoryPwd, enterPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(memoryPwd), []byte(enterPwd))
	return err == nil
}
