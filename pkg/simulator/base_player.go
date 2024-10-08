package simulator

import (
	"math/rand"

	"github.com/furudenipa/diceraceDP/config"
)

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

func (bp *BasePlayer) IsEnd() bool {
	return bp.remainingRolls == 0 && bp.remainingTickets[0] == 0 && bp.remainingTickets[1] == 0 && bp.remainingTickets[2] == 0 && bp.remainingTickets[3] == 0 && bp.remainingTickets[4] == 0 && bp.remainingTickets[5] == 0
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

// Clone returns a deep copy of the player
// Clone id made by NewBasePlayer(initRemainingRollsCopy, initRemainingTicketsCopy, initSquareCopy)
func (bp *BasePlayer) Clone() *BasePlayer {
	initRemainingRollsCopy := bp.initRemainingRolls
	initRemainingTicketsCopy := make([]int, len(bp.initRemainingTickets))
	copy(initRemainingTicketsCopy, bp.initRemainingTickets)
	initSquareCopy := bp.initSquare
	return NewBasePlayer(initRemainingRollsCopy, initRemainingTicketsCopy, initSquareCopy)
}

func NewBasePlayer(remainingRolls int, remainingTickets []int, square int) *BasePlayer {
	ticketsCopy := make([]int, len(remainingTickets))
	copy(ticketsCopy, remainingTickets)

	return &BasePlayer{
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

// return remainingRolls, remainingTickets, square
func (bp *BasePlayer) GetState() (int, []int, int) {
	return bp.remainingRolls, bp.remainingTickets, bp.square
}

func (bp *BasePlayer) GetTotalReward() float64 {
	return bp.totalReward
}

func (bp *BasePlayer) GetTotalItems() map[string]int {
	return bp.totalItems
}
