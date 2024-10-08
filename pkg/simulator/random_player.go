package simulator

import (
	"math/rand"

	"github.com/furudenipa/diceraceDP/config"
)

type RandomPlayer struct {
	*BasePlayer
}

func (rp *RandomPlayer) TakeAction() (byte, bool) {
	if rp.IsEnd() {
		return 7, false
	}

	actions := make([]byte, 0)
	if rp.remainingRolls > 0 {
		actions = append(actions, config.ActionRoll)
	}
	for i := 0; i < 6; i++ {
		if rp.remainingTickets[i] > 0 {
			actions = append(actions, byte(i+1))
		}
	}

	action := actions[rand.Intn(len(actions))]
	rp.move(action)
	return action, true
}

func (rp *RandomPlayer) Clone() Player {
	return NewRandomPlayer(rp.BasePlayer.Clone())
}

func NewRandomPlayer(basePlayer *BasePlayer) *RandomPlayer {
	return &RandomPlayer{
		BasePlayer: basePlayer,
	}
}
