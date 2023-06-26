package main

import (
	"github.com/name5566/leaf"
	"leaf_game/game"
	"leaf_game/gate"
	"leaf_game/login"
)

func main() {
	leaf.Run(
		game.Module,
		gate.Module,
		login.Module,
	)
}
