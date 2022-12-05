package utils

import (
	"GinRESTful/restapi/global"
	"golang.org/x/crypto/bcrypt"
)

// SetPassword 设置密码，加密密码
func SetPassword(password string) ([]byte, error) {
	userInfo := global.Settings.UserInfo
	pwdBytes, err := bcrypt.GenerateFromPassword([]byte(password), userInfo.PwdEncDiff)
	if err != nil {
		return nil, err
	}
	return pwdBytes, nil
}

// CheckPassword 校验密码
func CheckPassword(enterPwd, memoryPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(enterPwd), []byte(memoryPwd))
	return err == nil
}
