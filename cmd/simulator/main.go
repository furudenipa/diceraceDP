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

// TODO: メモリ使用量がbinの2倍に膨れる メモリプロファイル解析
func main() {
	iterate := 10000000
	numWorkers := runtime.NumCPU()

	players := createPlayers()
	names := []string{"Random", "Custom", "Policy"}

	for idx, player := range players {
		runSimulation(player, names[idx], iterate, numWorkers)
	}
}

func createPlayers() []simulator.Player {
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
		&strides,
		policyPointer,
	)
	return []simulator.Player{player1, player2, player3}
}

func runWorker(player simulator.Player, iterations int, rewardsChan chan<- float64, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < iterations; i++ {
		reward, _ := simulator.PlayOut(player)
		rewardsChan <- reward
	}
}

func runSimulation(player simulator.Player, name string, iterations int, numWorkers int) {
	s := time.Now()

	rewardsChan := make(chan float64, numWorkers*100)
	var wg sync.WaitGroup
	stats := stats.Stats{}

	go func() {
		for reward := range rewardsChan {
			stats.Add(reward)
		}
	}()

	simulationsPerWorker := iterations / numWorkers
	remainingSimulations := iterations % numWorkers

	for i := 0; i < numWorkers; i++ {
		sims := simulationsPerWorker
		if i == numWorkers-1 {
			sims += remainingSimulations
		}
		wg.Add(1)
		go runWorker(player.Clone(), sims, rewardsChan, &wg)
	}

	wg.Wait()
	close(rewardsChan)

	fmt.Println("-----------------------------Player Statistics------------------------------------")
	fmt.Printf("%-15s %-15s %-15s %-15s %-15s\n", "PlayerName", "Mean Reward", "Variance", "Std Dev", "Time")
	fmt.Printf("%-15s %-15.2f %-15.2f %-15.2f %-15s\n", name, stats.Mean(), stats.Variance(), stats.StdDev(), time.Since(s))
	fmt.Println("----------------------------------------------------------------------------------")
}
