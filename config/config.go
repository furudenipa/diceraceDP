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

func GetConfigPath() string {
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

// テストの初期化時に設定をリセットするためのセットアップ関数
func SetTestConfig() {
	os.Setenv("DICERACE_CONFIG", "test")
	configPath := GetConfigPath()
	SetConfig(
		filepath.Join(configPath, "items.yaml"),
		filepath.Join(configPath, "cells.yaml"),
	)
}

func init() {
	dir, err := os.Getwd()
	if err != nil {
		slog.Error("Error getting current working directory: " + err.Error())
		os.Exit(1)
	}
	slog.Info("Current working directory: " + dir)

	configPath := GetConfigPath()
	SetConfig(
		filepath.Join(configPath, "items.yaml"),
		filepath.Join(configPath, "cells.yaml"),
	)
}
