package config

import (
	"fmt"
	"log/slog"
	"os"
	"slices"

	"gopkg.in/yaml.v2"
)

// Items構造体はitems.yamlの内容を表します
type Items struct {
	Items map[string]Item `yaml:"items"`
}

// Item構造体は各アイテムの詳細を表します
type Item struct {
	Reward float64 `yaml:"reward"`
}

// Cells構造体はcells.yamlの内容を表します
type Cells struct {
	Cells []Cell `yaml:"cells"`
}

// Cell構造体は各セルの詳細を表します
type Cell struct {
	Item  string `yaml:"item"`
	Count int    `yaml:"count"`
}

func ReadYaml(filepath string, data interface{}) {
	dataBytes, err := os.ReadFile(filepath)
	if err != nil {
		slog.Error("%sの読み込み中にエラーが発生しました: %v", filepath, err)
	}
	err = yaml.Unmarshal(dataBytes, data)
	if err != nil {
		slog.Error("%sのパース中にエラーが発生しました: %v", filepath, err)
	}
}

func configValidation(items *Items, cells *Cells) error {
	var itemNames []string

	// items.yamlの検証
	for name, item := range items.Items {
		itemNames = append(itemNames, name)
		if item.Reward < 0 {
			return fmt.Errorf("items.yaml: %sの報酬が負の値です: %.2f", name, item.Reward)
		}
	}

	// cells.yamlの検証
	for _, cell := range cells.Cells {
		if cell.Count < 0 {
			return fmt.Errorf("cells.yaml: %sの報酬が負の値です: %d", cell.Item, cell.Count)
		}
		if !slices.Contains(itemNames, cell.Item) {
			return fmt.Errorf("cells.yaml: %sがitems.yamlに存在しません", cell.Item)
		}
	}

	// セルの数がNumSquaresと一致するか検証
	if len(cells.Cells) != NumSquares {
		return fmt.Errorf("cells.yaml: セルの数がNumSquaresと一致しません: %d != %d", len(cells.Cells), NumSquares)
	}

	return nil
}

func LoadConfig(itemsPath, cellsPath string) (*Items, *Cells) {
	// items.yamlを読み込む
	if itemsPath == "" {
		itemsPath = "../../config/yaml/dev/items.yaml"
	}
	var items Items
	ReadYaml(itemsPath, &items)

	// cells.yamlを読み込む
	if cellsPath == "" {
		cellsPath = "../../config/yaml/dev/cells.yaml"
	}
	var cells Cells
	ReadYaml(cellsPath, &cells)

	if err := configValidation(&items, &cells); err != nil {
		slog.Error("設定ファイルの検証中にエラーが発生しました: %v", slog.String("error", err.Error()))
		os.Exit(1)
	}

	return &items, &cells
}

func makeRewards(items *Items, cells *Cells) *[]float64 {
	rewards := make([]float64, NumSquares)
	for i, cell := range cells.Cells {
		rewards[i] = float64(cell.Count) * items.Items[cell.Item].Reward
	}
	return &rewards
}

func makeExpRewards(rewards *[]float64) *[]float64 {
	expRewards := make([]float64, NumSquares)
	var sum float64
	for i := 0; i < NumSquares; i++ {
		sum = 0
		for j := 1; j <= 6; j++ {
			sum += (*rewards)[(i+j)%NumSquares]
		}
		expRewards[i] = sum
	}
	return &expRewards
}
