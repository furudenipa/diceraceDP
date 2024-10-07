package simulator

import (
	"math/rand"

	"github.com/furudenipa/diceraceDP/config"
	"github.com/furudenipa/diceraceDP/reader"
)

type Player interface {
	TakeAction()
}

type BasePlayer struct {
	remainingRolls       int
	remainingTickets     []int
	square               int
	totalReward          float64
	totalItems           map[string]int
	initRemainingRolls   int
	initRemainingTickets []int
	initSquare           int
}

type RandomPlayer struct {
	BasePlayer
}

type AiPlayer struct {
	BasePlayer
	strides []int
	policy  []byte // Example policy to determine actions
}

func (ap *AiPlayer) TakeAction() (byte, bool) {
	if ap.remainingRolls == 0 && ap.remainingTickets[0] == 0 && ap.remainingTickets[1] == 0 && ap.remainingTickets[2] == 0 && ap.remainingTickets[3] == 0 && ap.remainingTickets[4] == 0 && ap.remainingTickets[5] == 0 {
		return 0, false
	}
	idx := reader.GetFlatIndex(ap.remainingRolls, ap.square, ap.remainingTickets, ap.strides)
	action := ap.policy[idx] % 8
	ap.move(action)
	return action, true
}

func (bp *BasePlayer) move(action byte) {
	switch action {
	case config.ActionRoll:
		bp.remainingRolls--
		pips := rand.Intn(6) + 1
		bp.square = (bp.square + pips) % config.NumSquares
	case config.ActionTicket1:
		bp.remainingTickets[0]--
		bp.square = (bp.square + 1) % config.NumSquares
	case config.ActionTicket2:
		bp.remainingTickets[1]--
		bp.square = (bp.square + 2) % config.NumSquares
	case config.ActionTicket3:
		bp.remainingTickets[2]--
		bp.square = (bp.square + 3) % config.NumSquares
	case config.ActionTicket4:
		bp.remainingTickets[3]--
		bp.square = (bp.square + 4) % config.NumSquares
	case config.ActionTicket5:
		bp.remainingTickets[4]--
		bp.square = (bp.square + 5) % config.NumSquares
	case config.ActionTicket6:
		bp.remainingTickets[5]--
		bp.square = (bp.square + 6) % config.NumSquares
	}
	if bp.square == config.AdvanceSixSquare {
		bp.square = (bp.square + 6) % config.NumSquares
	}
	if bp.square == config.TicketSquare {
		i := rand.Intn(6)
		bp.remainingTickets[i] = min(bp.remainingTickets[i]+1, config.MaxTickets-1)
	}

	item := config.ItemsList[bp.square]
	bp.totalReward += item.Reward
	bp.totalItems[item.Name] += item.Count
}

func (bp *BasePlayer) Reset() {
	ticketsCopy := make([]int, len(bp.initRemainingTickets))
	copy(ticketsCopy, bp.initRemainingTickets)

	bp.remainingRolls = bp.initRemainingRolls
	bp.remainingTickets = ticketsCopy
	bp.square = bp.initSquare
	bp.totalReward = 0
	bp.totalItems = make(map[string]int)
}

func NewBasePlayer(remainingRolls int, remainingTickets []int, square int) BasePlayer {
	ticketsCopy := make([]int, len(remainingTickets))
	copy(ticketsCopy, remainingTickets)

	return BasePlayer{
		remainingRolls:       remainingRolls,
		remainingTickets:     ticketsCopy,
		square:               square,
		initRemainingRolls:   remainingRolls,
		initRemainingTickets: remainingTickets,
		initSquare:           square,
		totalReward:          0,
		totalItems:           make(map[string]int),
	}
}

func NewAiPlayer(basePlayer BasePlayer, strides []int, policy []byte) AiPlayer {
	return AiPlayer{
		BasePlayer: basePlayer,
		strides:    strides,
		policy:     policy,
	}
}

func (bp *BasePlayer) GetState() (int, []int, int) {
	return bp.remainingRolls, bp.remainingTickets, bp.square
}

func (bp *BasePlayer) GetTotalReward() float64 {
	return bp.totalReward
}

func (bp *BasePlayer) GetTotalItems() map[string]int {
	return bp.totalItems
}
