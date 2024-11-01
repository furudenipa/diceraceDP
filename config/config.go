package config

import (
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
)

const (
	NumSquares       = 18
	MaxTickets       = 10
	MaxRolls         = 100
	TicketSquare     = 0
	AdvanceSixSquare = 12
)

const NumDimensions = 8

const (
	// Roll and Ticket
	ActionRoll = iota
	ActionTicket1
	ActionTicket2
	ActionTicket3
	ActionTicket4
	ActionTicket5
	ActionTicket6
	ActionNothing
)

var I Items
var C Cells
var Rewards []float64
var ExpRewards []float64

func getConfigPath() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		slog.Error("Error getting current file path")
		os.Exit(1)
	}
	dir := filepath.Dir(filename)
	projectRoot := filepath.Join(dir, "../")
	env := os.Getenv("DICERACE_CONFIG")

	var configDir string
	switch env {
	case "test":
		configDir = "config/yaml/test"
	default:
		configDir = "config/yaml/dev"
	}

	configPath := filepath.Join(projectRoot, configDir)
	return configPath
}

func SetConfig(itemsPath, cellsPath string) {
	tmpI, tmpC := LoadConfig(itemsPath, cellsPath)
	I = *tmpI
	C = *tmpC
	Rewards = *makeRewards(tmpI, tmpC)
	ExpRewards = *makeExpRewards(&Rewards)
}

func init() {
	dir, err := os.Getwd()
	if err != nil {
		slog.Error("Error getting current working directory: " + err.Error())
		os.Exit(1)
	}
	slog.Info("Current working directory: " + dir)

	configPath := getConfigPath()
	SetConfig(
		filepath.Join(configPath, "items.yaml"),
		filepath.Join(configPath, "cells.yaml"),
	)
}

// type ItemInfo struct {
// 	Reward float64 `yaml:"reward"`
// }

// type Config struct {
// 	Items map[ItemType]ItemInfo `yaml:"items"`
// 	Cells []Cell                `yaml:"cells"`
// }

// func LoadConfig(filename string) (*Config, error) {
// 	data, err := ioutil.ReadFile(filename)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var config Config
// 	err = yaml.Unmarshal(data, &config)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// リワードの計算
// 	for idx := range config.Cells {
// 		cell := &config.Cells[idx]
// 		itemInfo, ok := config.Items[cell.Item]
// 		if !ok {
// 			return nil, fmt.Errorf("unknown item type: %s", cell.Item)
// 		}
// 		cell.Reward = itemInfo.Reward * float64(cell.Count)
// 	}

// 	return &config, nil
// }

// type ItemType int

// const (
// 	None ItemType = iota
// 	Stone
// 	Credit
// 	Report
// 	Miyu
// 	Eleph
// )

// type ItemInfo struct {
//     Name   string
//     Reward float64
// }

// var ItemMap = map[ItemType]ItemInfo{
// 	None:   {Name= "none", Reward = 0},
// 	Stone:  {Name= "stone", Reward = 4.3},
// 	Credit: {Name= "credit", Reward = 13.5},
// 	Report: {Name= "report", Reward = 4	},
// 	Miyu:   {Name= "miyu", Reward = 0},
// 	Eleph:  {Name= "eleph", Reward =50},
// }

// type Cell struct {
//     Item     ItemType
//     Count    int
//     Reward   float64
// }

// var Cells []Cell

// func init() {
//     Cells = []Cell{
//         {Item: None, Count: 0, Reward: 0},
//         {Item: Stone, Count: 10, Reward: RewardMap[Stone] * 10},
//         {Item: Credit, Count: 20, Reward: RewardMap[Credit] * 20},
//         {Item: Report, Count: 17, Reward: RewardMap[Report] * 17},
//         {Item: Miyu, Count: 5, Reward: RewardMap[Miyu] * 5},
//         {Item: Credit, Count: 32, Reward: RewardMap[Credit] * 32},
//         {Item: Stone, Count: 15, Reward: RewardMap[Stone] * 15},
// 		{Item: Report, Count: 22, Reward: RewardMap[Report] * 22},
// 		{Item: Eleph, Count: 8, Reward: RewardMap[Eleph] * 8},
// 		{Item: Miyu, Count: 7, Reward: RewardMap[Miyu] * 7},
// 		{Item: Eleph, Count: 12, Reward: RewardMap[Eleph] * 12},
// 		{Item: Credit, Count: 24, Reward: RewardMap[Credit] * 24},
// 		{Item: None, Count: 0, Reward: 0},
// 		{Item: Credit, Count: 16, Reward: RewardMap[Credit] * 16},
// 		{Item: Stone, Count: 12, Reward: RewardMap[Stone] * 12},
// 		{Item: Report, Count: 10, Reward: RewardMap[Report] * 10},
// 		{Item: Eleph, Count: 6, Reward: RewardMap[Eleph] * 6},
// 		{Item: Miyu, Count: 4, Reward: RewardMap[Miyu] * 4},
// 	}
// }
