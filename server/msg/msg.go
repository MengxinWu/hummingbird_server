package msg

import (
	"github.com/name5566/leaf/network/json"
)

var (
	// Processor 使用json消息处理器
	Processor = json.NewProcessor()
	Num       = int32(1)
)

func init() {
	// 注册protobuf消息Hello
	Processor.Register(&Hello{})
}

// Hello 定义JSON格式消息Hello
type Hello struct {
	Name string
	Num  int32
}
