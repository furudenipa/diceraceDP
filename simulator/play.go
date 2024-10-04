package simulator

type Player interface {
	TakeAction()
}

type AiPlayer struct {
	remainingRolls   int
	remainingTickets []int
	square           int
	totalReward      float64
	totalItems       []int
	policy           []byte
	strides          []int
}

func (a *AiPlayer) TakeAction() {

}
