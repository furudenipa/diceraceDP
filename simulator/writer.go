package simulator

/*
import (
	"encoding/binary"
	"os"

	"github.com/furudenipa/diceraceDP/config"
)

func writer(stateValues *[config.NumSteps][config.NumSquares][config.MaxTickets][config.MaxTickets][config.MaxTickets][config.MaxTickets][config.MaxTickets][config.MaxTickets]float64) {

	// ファイルに書き込む
	file, err := os.Create("floats6d.bin")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// すべての要素をバイナリに書き込む
	for i := 0; i < config.NumSteps; i++ {
		for j := 0; j < config.NumSquares; j++ {
			for k := 0; k < config.MaxTickets; k++ {
				for l := 0; l < config.MaxTickets; l++ {
					for m := 0; m < config.MaxTickets; m++ {
						for n := 0; n < config.MaxTickets; n++ {
							for o := 0; o < config.MaxTickets; o++ {
								for p := 0; p < config.MaxTickets; p++ {
									// 仮のデータを使用
									value := float64(stateValues[i][j][k][l][m][n][o][p])
									// バイナリ形式で書き込み
									err := binary.Write(file, binary.LittleEndian, value)
									if err != nil {
										panic(err)
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

func miniWriter() {
	const numSteps = 5   // 例として小さな値を設定
	const numSquares = 5 // 例として小さな値を設定
	const maxTickets = 5 // 例として小さな値を設定

	// ファイルを作成
	file, err := os.Create("mini.bin")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 8次元配列をバイナリデータとしてエンコードしファイルに書き込む
	for i := 0; i < numSteps; i++ {
		for j := 0; j < numSquares; j++ {
			for k := 0; k < maxTickets; k++ {
				for l := 0; l < maxTickets; l++ {
					for m := 0; m < maxTickets; m++ {
						for n := 0; n < maxTickets; n++ {
							for o := 0; o < maxTickets; o++ {
								for p := 0; p < maxTickets; p++ {
									// 仮のデータを使用
									value := float64(i + j + k + l + m + n + o + p)
									// バイナリ形式で書き込み
									err := binary.Write(file, binary.LittleEndian, value)
									if err != nil {
										panic(err)
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
*/
