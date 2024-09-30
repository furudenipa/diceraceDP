package simulator

type Player interface {
	TakeAction()
}

type AiPlayer struct {
	step        int
	square      int
	tickets     []int
	totalReward float64
	totalItems  []int
	dummyPolicy []byte
	// policy
}

func (a *AiPlayer) TakeAction() {

}
