package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5 计算字符串的MD5值，参考链接：https://wangbjun.site/2020/coding/golang/md5.html
// 参数：
//		str：需要MD5的字符串
// 返回值：
//		string：MD5之后的字符串
func MD5(str string) string {
	sum := md5.Sum([]byte(str))
	return hex.EncodeToString(sum[:])
}
