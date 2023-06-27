package gate

import (
	"server/game"
	"server/msg"
)

func init() {
	// 这里指定消息Hello路由到模块game
	// 模块间使用chanRPC通讯 消息路由也使用chanRPC
	msg.Processor.SetRouter(&msg.Hello{}, game.ChanRPC)
}
