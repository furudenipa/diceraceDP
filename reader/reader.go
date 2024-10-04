package reader

import (
	"io"
	"log"
	"os"

	"github.com/furudenipa/diceraceDP/config"
)

func PolicyReader(filePath string) *[]byte {

	// ファイルを開く
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("ファイルを開く際にエラーが発生しました: %v", err)
	}
	defer file.Close()

	// ファイル情報を取得
	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatalf("ファイル情報を取得できませんでした: %v", err)
	}

	// 期待されるデータサイズを計算
	expectedSize := config.NumSteps * config.NumSquares * Pow(config.MaxTickets, 6)
	if int64(expectedSize) != fileInfo.Size() {
		log.Fatalf("ファイルサイズが期待されるサイズ (%d) と一致しません。実際のサイズ: %d", expectedSize, fileInfo.Size())
	}

	// データを一度に読み込む
	data := make([]byte, expectedSize)
	_, err = io.ReadFull(file, data)
	if err != nil {
		log.Fatalf("ファイルを読み込む際にエラーが発生しました: %v", err)
	}

	return &data
}
