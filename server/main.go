package main

import (
	"log"

	"server/game"
	"server/gate"
	"server/login"

	"github.com/name5566/leaf"
	lconf "github.com/name5566/leaf/conf"
)

func main() {
	lconf.LogLevel = "debug"
	lconf.LogFlag = log.LstdFlags

	leaf.Run(
		game.Module,
		gate.Module,
		login.Module,
	)
}