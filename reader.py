import numpy as np

def load_array_from_binary(file_path):
    # ファイルを読み込み、float64型としてデータをロード
    data = np.fromfile(file_path, dtype=np.float64)

    # 8次元配列の形状を指定
    numSteps = 5
    numSquares = 5
    maxTickets = 5

    # データを8次元配列に再形成
    reshaped_data = data.reshape((numSteps, numSquares, maxTickets, maxTickets, maxTickets, maxTickets, maxTickets, maxTickets))
    return reshaped_data

# ファイルパスを指定
file_path = 'mini.bin'

# 配列を読み込み
array = load_array_from_binary(file_path)

# 読み込んだ配列の内容を表示
print(array)
