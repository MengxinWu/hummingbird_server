package internal

import (
	"fmt"
	"reflect"
	"sync/atomic"

	"server/msg"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

func init() {
	// 向game模块注册Hello消息处理函数handleHello
	handler(&msg.Hello{}, handleHello)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleHello(args []interface{}) {
	// 收到的Test消息
	m := args[0].(*msg.Hello)
	// 消息的发送者
	a := args[1].(gate.Agent)
	// 输出收到的消息的内容
	fmt.Println("hello")
	log.Debug("hello %v", m.Name)
	// 给发送者回应Test消息
	atomic.AddInt32(&msg.Num, 1)
	a.WriteMsg(&msg.Hello{
		Name: "client",
		Num:  msg.Num,
	})
}
