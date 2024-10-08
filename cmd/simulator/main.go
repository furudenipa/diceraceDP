package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/furudenipa/diceraceDP/pkg/reader"
	"github.com/furudenipa/diceraceDP/pkg/simulator"
	"github.com/furudenipa/diceraceDP/pkg/stats"
)

func main() {
	iterate := 10000000
	numWorkers := runtime.NumCPU()

	strides := reader.ComputeStrides()
	policyPointer := reader.ReadPolicy("../../data/policy3.bin")

	player1 := simulator.NewRandomPlayer(
		simulator.NewBasePlayer(
			99,
			[]int{1, 1, 1, 1, 1, 1},
			0,
		),
	)
	player2 := simulator.NewCustomPlayer(
		simulator.NewBasePlayer(
			99,
			[]int{1, 1, 1, 1, 1, 1},
			0,
		),
		simulator.SampleStrategy1,
	)
	player3 := simulator.NewAiPlayer(
		simulator.NewBasePlayer(
			99,
			[]int{1, 1, 1, 1, 1, 1},
			0,
		),
		strides,
		policyPointer,
	)

	players := []simulator.Player{player1, player2, player3}
	names := []string{"Random", "Custom", "Policy"}

	for idx, player := range players {
		s := time.Now()

		rewardsChan := make(chan float64, numWorkers*100)
		var wg sync.WaitGroup
		stats := stats.Stats{}

		// consume rewards from the channel
		go func() {
			for reward := range rewardsChan {
				stats.Add(reward)
			}
		}()

		// worker function
		worker := func(player simulator.Player, iterations int) {
			defer wg.Done()
			for i := 0; i < iterations; i++ {
				reward, _ := simulator.RunSimulation(player)
				rewardsChan <- reward
			}
		}

		simulationsPerWorker := iterate / numWorkers
		remainingSimulations := iterate % numWorkers
		fmt.Printf("Running %d simulations using %d workers\n", iterate, numWorkers)
		for i := 0; i < numWorkers; i++ {
			sims := simulationsPerWorker
			if i == numWorkers-1 {
				sims += remainingSimulations
			}
			wg.Add(1)
			go worker(player.Clone(), sims)
		}

		go func() {
			wg.Wait()
			close(rewardsChan)
		}()

		wg.Wait()

		// 集計結果を表示
		fmt.Println()
		fmt.Printf("PlayerName: %s\n", names[idx])
		fmt.Printf(" Mean Reward: %f\n", stats.Mean())
		fmt.Printf(" Variance: %f\n", stats.Variance())
		fmt.Printf(" Standard Deviation: %f\n", stats.StdDev())
		fmt.Printf(" Time: %s\n", time.Since(s))
	}
}

/*
func _hoge() {

	fmt.Println("--------Start simulation--------")
	for n, player := range players {
		avgRewrad := 0.0
		s := time.Now()
		for i := 0; i < iterate; i++ {
			player.Reset()
			for {
				_, ok := player.TakeAction()
				if !ok {
					break
				}
			}
			avgRewrad += player.GetTotalReward()
		}
		avgRewrad /= float64(iterate)
		fmt.Println("PlayerName:", names[n], "  Average reward:", avgRewrad)
		fmt.Println(" Time:", time.Since(s))
	}
	fmt.Println("----------End simulation--------")
}
*/
