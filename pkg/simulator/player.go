package simulator

type Player interface {
	TakeAction() (byte, bool)
	Reset()
	GetState() (int, []int, int)
	GetTotalReward() float64
	GetTotalItems() map[string]int
	Clone() Player
}

func RunSimulation(player Player) (float64, map[string]int) {
	player.Reset()
	for {
		_, ok := player.TakeAction()
		if !ok {
			break
		}
	}
	return player.GetTotalReward(), player.GetTotalItems()
}
