package reader

import (
	"io"
	"log/slog"
	"os"

	"github.com/furudenipa/diceraceDP/config"
)

// 指定されたファイルからpolicyを読み込む。
// configで設定されているサイズと一致する必要があります。
func ReadPolicy(filePath string) *[]byte {
	// ファイルを開く
	file, err := os.Open(filePath)
	if err != nil {
		slog.Error("Failed to open file", slog.String("error", err.Error()))
		os.Exit(1)
	}
	defer file.Close()

	// ファイル情報を取得
	fileInfo, err := file.Stat()
	if err != nil {
		slog.Error("Failed to get file information", slog.String("error", err.Error()))
		os.Exit(1)
	}

	// 期待されるデータサイズを計算
	expectedSize := config.MaxRolls * config.NumSquares * Pow(config.MaxTickets, 6)
	if int64(expectedSize) != fileInfo.Size() {
		slog.Error("File size does not match the expected size", slog.Int("expectedSize", expectedSize), slog.Int64("actualSize", fileInfo.Size()))
		os.Exit(1)
	}

	// データを一度に読み込む
	data := make([]byte, expectedSize)
	_, err = io.ReadFull(file, data)
	if err != nil {
		slog.Error("Failed to read file", slog.String("error", err.Error()))
	}

	return &data
}
