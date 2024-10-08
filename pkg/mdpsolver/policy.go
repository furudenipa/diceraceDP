package mdpsolver

import "github.com/furudenipa/diceraceDP/config"

type policy struct {
	policy [config.NumSquares][config.MaxTickets][config.MaxTickets][config.MaxTickets][config.MaxTickets][config.MaxTickets][config.MaxTickets]byte
}

func (p *policy) GetAction(square int, remainingTickets []int) byte {
	return p.policy[square][remainingTickets[0]][remainingTickets[1]][remainingTickets[2]][remainingTickets[3]][remainingTickets[4]][remainingTickets[5]]
}

func (p *policy) SetAction(square int, remainingTickets []int, action byte) {
	p.policy[square][remainingTickets[0]][remainingTickets[1]][remainingTickets[2]][remainingTickets[3]][remainingTickets[4]][remainingTickets[5]] = action
}

func (p *policy) GetPolicyTwoDimensionalSlice(i, j, k, l, m int) [10][10]byte {
	return p.policy[i][j][k][l][m]
}
