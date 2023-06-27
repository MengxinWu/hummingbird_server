package login

import "server/login/internal"

var (
	// Module 实例化模块login
	Module = new(internal.Module)
	// ChanRPC 暴露ChanRPC
	ChanRPC = internal.ChanRPC
)
