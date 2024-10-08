package simulator

import (
	"log"
	"math/rand"

	"github.com/furudenipa/diceraceDP/config"
)

type CustomPlayer struct {
	*BasePlayer
	strategy func(int, []int, int) byte
}

func (cp *CustomPlayer) TakeAction() (byte, bool) {
	if cp.IsEnd() {
		return 7, false
	}

	action := cp.strategy(cp.remainingRolls, cp.remainingTickets, cp.square)
	if action == config.ActionNothing {
		log.Fatal("Invalid action")
		return 7, false
	}
	cp.move(action)
	return action, true
}

func (cp *CustomPlayer) Clone() Player {
	return NewCustomPlayer(cp.BasePlayer.Clone(), cp.strategy)
}

func NewCustomPlayer(basePlayer *BasePlayer, strategy func(int, []int, int) byte) *CustomPlayer {
	return &CustomPlayer{
		BasePlayer: basePlayer,
		strategy:   strategy,
	}
}

// SampleStrategy1
// ヴァッシュTV (https://www.youtube.com/watch?v=TO21AHyX008) による戦略
func SampleStrategy1(remainingRolls int, remainingTickets []int, square int) byte {
	// vashTV strategy
	switch square {
	case 2, 4:
		if remainingTickets[5] > 0 {
			return config.ActionTicket6
		}
	case 5:
		if remainingTickets[4] > 0 {
			return config.ActionTicket5
		}
	case 8, 10:
		if remainingTickets[1] > 0 {
			return config.ActionTicket2
		}
	case 14:
		if remainingTickets[3] > 0 {
			return config.ActionTicket4
		}
	case 15:
		if remainingTickets[2] > 0 {
			return config.ActionTicket3
		}
	case 17:
		if remainingTickets[0] > 0 {
			return config.ActionTicket1
		}
	}

	// Roll dice if remainingRolls > 0
	if remainingRolls > 0 {
		return config.ActionRoll
	}

	// Other situations.
	actions := make([]byte, 0)
	for i := 0; i < 6; i++ {
		if remainingTickets[i] > 0 {
			actions = append(actions, byte(i+1))
		}
	}
	action := actions[rand.Intn(len(actions))]
	return action
}
