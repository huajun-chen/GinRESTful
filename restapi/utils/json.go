package utils

import (
	"encoding/json"
	"io/ioutil"
)

// ReadJSON 读取json文件
// 参数：
//		filePath：文件的具体位置与文件名的组合，示例：xxx/xxx.json
// 返回值：
//		map[string]string：将json解析为map的结果
//		error：错误信息
func ReadJSON(filePath string) (map[string]string, error) {
	// 读取json文件
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	// 解析json文件
	var translations map[string]string
	if err := json.Unmarshal(data, &translations); err != nil {
		return nil, err
	}
	return translations, nil
}
