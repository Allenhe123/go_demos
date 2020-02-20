package search

import (
	"encoding/json"
	"os"
)

const dataFile = "data/data.json"

// Feed 包含我们需要处理的数据源的信息
// `引号里的部分被称作标记（tag）。这个标记里描述了JSON解码的元数据，
// 用于创建 Feed 类型值的切片。每个标记将结构类型里字段对应到JSON文档里指定名字的字段。
type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

// RetrieveFeeds 读取并反序列化源数据文件
func RetrieveFeeds() ([]*Feed, error) {
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}

	// 当函数返回时关闭文件
	defer file.Close()

	// 将文件解码到一个切片里
	// 这个切片的每一项是一个指向一个 Feed 类型值的指针
	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)
	// 这个函数不需要检查错误，调用者会做这件事
	return feeds, err
}
