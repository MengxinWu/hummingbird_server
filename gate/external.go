package gate

import "leaf_game/gate/internal"

var (
	// 实例化模块
	Module = new(internal.Module)
	// 暴露 ChanRPC
	ChanRPC = internal.ChanRPC
)
