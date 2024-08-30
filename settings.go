package main

const (
	numSquares        = 18
	maxTickets         = 10
	numSteps           = 100
	ticketSquare      = 0
	advanceSixSquare = 12
	STONE              = 1
	CREDIT             = 3
	REPORT             = 2
	MIYU               = 0
	ELEPH              = 5
)

/*
	var rewardsMap = [18]int{
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
*/
var rewardsMap = [18]int{
	0,
	10,
	20,
	30,
	40,
	50,
	60,
	70,
	80,
	90,
	100,
	110,
	120,
	130,
	140,
	150,
	160,
	170,
}

var diceRewardsMap [18]float64

func init() {
	sum := 0
	for i := 0; i < 18; i++ {
		sum = 0
		for j := 1; j <= 6; j++ {
			sum += rewardsMap[(i+j)%18]
		}
		diceRewardsMap[i] = float64(sum)
	}
}
