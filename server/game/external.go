package game

import "server/game/internal"

var (
	// Module 实例化模块game
	Module = new(internal.Module)
	// 暴露 ChanRPC
	ChanRPC = internal.ChanRPC
)
