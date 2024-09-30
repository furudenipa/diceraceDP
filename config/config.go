package config

const (
	NumSquares       = 18
	MaxTickets       = 10
	NumSteps         = 100
	TicketSquare     = 0
	AdvanceSixSquare = 12
	NumDimensions    = 8
)

const (
	// Rewards
	STONE  = 4.3  // yellow
	CREDIT = 13.5 // 100K
	REPORT = 4    // yellow
	MIYU   = 0
	ELEPH  = 50
)
const (
	numOfItemTypes = 5
)

var ItemsCount = [18]int{
	0, 10, 20, 17, 5, 32, 15, 22, 8, 7, 12, 24, 0, 16, 12, 10, 6, 4,
}

var RewardsMap = [18]float64{
	0,
	STONE * 10,
	CREDIT * 20,
	REPORT * 17,
	MIYU * 5,
	CREDIT * 32,
	STONE * 15,
	REPORT * 22,
	ELEPH * 8,
	MIYU * 7,
	ELEPH * 12,
	CREDIT * 24,
	0,
	CREDIT * 16,
	STONE * 12,
	REPORT * 10,
	ELEPH * 6,
	MIYU * 4,
}

var DiceRewardsMap [18]float64

func init() {
	var sum float64
	for i := 0; i < 18; i++ {
		sum = 0
		for j := 1; j <= 6; j++ {
			sum += RewardsMap[(i+j)%18]
		}
		DiceRewardsMap[i] = float64(sum)
	}
}
