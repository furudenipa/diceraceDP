package main

import (
	"fmt"

	"github.com/furudenipa/diceraceDP/reader"
	"github.com/furudenipa/diceraceDP/simulator"
)

func main() {
	//visualizer.Run("./data/policy3.bin")
	//mdpsolver.Run("data/policy_.bin")
	hoge()
}

func hoge() {
	aiPlayer := simulator.NewAiPlayer(
		simulator.NewBasePlayer(
			99,
			[]int{1, 1, 1, 1, 1, 1},
			0,
		),
		reader.ComputeStrides(),
		*reader.ReadPolicy("./data/policy3.bin"),
	)
	fmt.Println("SIMULATION START")
	cnt := 0
	avg := 0.0
	iterate := 1000000
	for i := 0; i < iterate; i++ {
		fmt.Println("-----Game ", i+1, "-----")

		aiPlayer.Reset()
		for {
			_, ok := aiPlayer.TakeAction()
			if !ok {
				break
			}
			//remainingRolls, remainingTickets, square := aiPlayer.GetState()
			//fmt.Println(" Action:", action)
			//fmt.Println("Rolls:", remainingRolls, " Tickets:", remainingTickets, " Square:", square)
			_, remainingTickets, _ := aiPlayer.GetState()
			if getmax(remainingTickets) == 9 {
				cnt++
			}
		}
		//fmt.Println("Total reward:", aiPlayer.GetTotalReward())
		//fmt.Println("Total items:", aiPlayer.GetTotalItems())
		avg += aiPlayer.GetTotalReward()
	}
	fmt.Println("SIMULATION END")
	fmt.Println("Average reward:", avg/float64(iterate))
	fmt.Println("max:", cnt)
}

func getmax(slice []int) int {
	max := slice[0]
	for _, v := range slice {
		if v > max {
			max = v
		}
	}
	return max
}
