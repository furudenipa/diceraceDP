import numpy as np

def load_array_from_binary(file_path):
    # ファイルを読み込み、float64型としてデータをロード
    data = np.fromfile(file_path, dtype=np.uint8)

    # 8次元配列の形状を指定
    numSteps = 100
    numSquares = 18
    maxTickets = 10

    # データを8次元配列に再形成
    reshaped_data = data.reshape((numSteps, numSquares, maxTickets, maxTickets, maxTickets, maxTickets, maxTickets, maxTickets))
    return reshaped_data

# ファイルパスを指定
file_path = 'policy.bin'

# 配列を読み込み
array = load_array_from_binary(file_path)

# 読み込んだ配列の内容を表示
while True:
    t1 = int(input('t1: '))
    t2 = int(input('t2: '))
    t3 = int(input('t3: '))
    t4 = int(input('t4: '))
    t5 = int(input('t5: '))
    t6 = int(input('t6: '))
    for i in range(0, 100, 10):
        print(array[i:i+10, :, t1, t2, t3, t4, t5, t6])
