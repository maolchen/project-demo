package utils

import "encoding/json"

// 结构体转map
func Struct2Map(obj interface{}) map[string]interface{} {
	jsonBytes, _ := json.Marshal(obj)
	m := make(map[string]interface{})
	json.Unmarshal(jsonBytes, &m)
	return m

}
