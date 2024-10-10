package simulator

import (
	"log/slog"

	"github.com/furudenipa/diceraceDP/pkg/reader"
)

type AiPlayer struct {
	*BasePlayer
	strides *[]int
	policy  *[]byte // Example policy to determine actions
}

func (ap *AiPlayer) TakeAction() (byte, bool) {
	if ap.IsEnd() {
		return 7, false
	}
	idx, err := reader.GetFlatIndex(ap.remainingRolls, ap.square, ap.remainingTickets, *ap.strides)
	if err != nil {
		slog.Error("Failed to get flat index", slog.String("error", err.Error()))
		return 7, false
	}
	policySlice := *ap.policy
	action := policySlice[idx] % 8
	ap.move(action)
	return action, true
}

// Clone returns a deep copy of the player
// Clone id made by NewAiPlayer(BasePlayer.Clone(), stridesCopy, policy)
func (ap *AiPlayer) Clone() Player {
	return &AiPlayer{
		BasePlayer: ap.BasePlayer.Clone(),
		strides:    ap.strides,
		policy:     ap.policy,
	}
}

func NewAiPlayer(basePlayer *BasePlayer, strides *[]int, policy *[]byte) *AiPlayer {
	return &AiPlayer{
		BasePlayer: basePlayer,
		strides:    strides,
		policy:     policy,
	}
}
