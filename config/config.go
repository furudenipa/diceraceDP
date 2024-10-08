package config

const (
	NumSquares       = 18
	MaxTickets       = 10
	MaxRolls         = 100
	TicketSquare     = 0
	AdvanceSixSquare = 12
	NumDimensions    = 8
)

const (
	// Rewards
	rewardStone  = 4.3  // yellow
	rewardCredit = 13.5 // 100K
	rewardReport = 4    // yellow
	rewardMiyu   = 0
	rewardEleph  = 50
)

const (
	// Roll and Ticket
	ActionRoll    = 0
	ActionTicket1 = 1
	ActionTicket2 = 2
	ActionTicket3 = 3
	ActionTicket4 = 4
	ActionTicket5 = 5
	ActionTicket6 = 6
	ActionNothing = 7
)

type Item struct {
	Name   string
	Count  int
	Reward float64
}

const (
	numOfItemTypes = 5
)

var ItemsCount = [18]int{
	0, 10, 20, 17, 5, 32, 15, 22, 8, 7, 12, 24, 0, 16, 12, 10, 6, 4,
}

var RewardsMap = [18]float64{
	0,
	rewardStone * 10,
	rewardCredit * 20,
	rewardReport * 17,
	rewardMiyu * 5,
	rewardCredit * 32,
	rewardStone * 15,
	rewardReport * 22,
	rewardEleph * 8,
	rewardMiyu * 7,
	rewardEleph * 12,
	rewardCredit * 24,
	0,
	rewardCredit * 16,
	rewardStone * 12,
	rewardReport * 10,
	rewardEleph * 6,
	rewardMiyu * 4,
}

var DiceRewardsMap [18]float64
var ItemsList [18]Item

func init() {

	ItemsList = [18]Item{
		{Name: "none", Count: 0, Reward: 0},                    //0
		{Name: "stone", Count: 10, Reward: rewardStone * 10},   // 1
		{Name: "credit", Count: 20, Reward: rewardCredit * 20}, //2
		{Name: "report", Count: 17, Reward: rewardReport * 17}, //3
		{Name: "miyu", Count: 5, Reward: rewardMiyu * 5},       //4
		{Name: "credit", Count: 32, Reward: rewardCredit * 32},
		{Name: "stone", Count: 15, Reward: rewardStone * 15},
		{Name: "report", Count: 22, Reward: rewardReport * 22},
		{Name: "eleph", Count: 8, Reward: rewardEleph * 8},
		{Name: "miyu", Count: 7, Reward: rewardMiyu * 7},
		{Name: "eleph", Count: 12, Reward: rewardEleph * 12},
		{Name: "credit", Count: 24, Reward: rewardCredit * 24},
		{Name: "none", Count: 0, Reward: 0},
		{Name: "credit", Count: 16, Reward: rewardCredit * 16},
		{Name: "stone", Count: 12, Reward: rewardStone * 12},
		{Name: "report", Count: 10, Reward: rewardReport * 10},
		{Name: "eleph", Count: 6, Reward: rewardEleph * 6},
		{Name: "miyu", Count: 4, Reward: rewardMiyu * 4},
	}

	var sum float64
	for i := 0; i < 18; i++ {
		sum = 0
		for j := 1; j <= 6; j++ {
			sum += RewardsMap[(i+j)%18]
		}
		DiceRewardsMap[i] = float64(sum)
	}
}
