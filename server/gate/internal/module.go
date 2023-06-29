package internal

import (
	"time"

	"server/game"
	"server/msg"

	"github.com/name5566/leaf/gate"
)

type Module struct {
	*gate.Gate
}

func (m *Module) OnInit() {
	m.Gate = &gate.Gate{
		MaxConnNum:      20000,
		PendingWriteNum: 2000,
		MaxMsgLen:       4096,
		WSAddr:          "0.0.0.0:3653",
		HTTPTimeout:     10 * time.Second,
		TCPAddr:         "0.0.0.0:3563",
		LenMsgLen:       2,
		LittleEndian:    false,
		Processor:       msg.Processor,
		AgentChanRPC:    game.ChanRPC,
	}
}
