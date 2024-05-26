package main

import (
	"fmt"
)

// コインの額面が日本円のように、1, 5, 10, 50, 100, 500と決められていたら、与えられた額に対して、額の大きいコインから引いて（割って）いけば最小枚数が求まる
// このような「そのときの（その時点の）最適な解（方法）を選んでいくアルゴリズム」を貪欲法（greedy method）という
func coinChangingProblem() {
	targetPrice := 15  // 支払う金額
	coinNum := 6  // 利用可能なコインの枚数
	coinValueList := [6]int{1, 2, 7, 8, 12, 50}  // 各コインの額（今回の問題文は既に昇順に並んでいたが、仮に並んでいない場合は並べる必要がありそう）
	minimumNumberList := make([][]int, coinNum)
	for i, _ := range minimumNumberList {
		minimumNumberList[i] = make([]int, targetPrice+1)
	}
	for coinIndex, coinValue := range coinValueList {
		// 各コインを使用したときの最小枚数を順に求めていく
		for price, _ := range minimumNumberList[coinIndex] {
			// price == 0 の場合、必要なコインの枚数は0である
			// minimumNumberList内の値は全て０で初期化されているので、再度代入する必要がない
			// price != 0 の場合だけ、最小必要枚数の更新が必要なので、そのような条件を設ける
			if price != 0 {
				// 問題文より、額が１のコインは必ず存在するという前提があるので、下記条件を問答無用で設定可能
				// 額が１より小さいコインは存在せず、そのコインを使用しなかった場合の最小枚数との比較ができない
				// また、額が１のコインを用いた場合、各priceを支払うための枚数はpriceになるため、それを下記コードにより対応する
				if coinValue == 1 {
					minimumNumberList[coinIndex][price] = price
				} else {  // コインの額が１ではない場合のみ、そのコインを使用した場合と使用しない場合で最小枚数の比較を行い、小さい枚数で更新する
					// 今回のコインを使用しなかった場合の最小枚数を求める
					priceInUnused := minimumNumberList[coinIndex-1][price]
					// coinValueは昇順に回ってくる
					// 当然、額が大きいコインをできるだけ使用した方が必要枚数は少なくなるので、priceをcoinValueで割って求められた値を現時点での必要最小枚数とする
					minimumNumber := price / coinValue
					remainder := price % coinValue
					minimumNumber += minimumNumberList[coinIndex-1][remainder]
					if priceInUnused <= minimumNumber {
						minimumNumberList[coinIndex][price] = priceInUnused
					} else {
						minimumNumberList[coinIndex][price] = minimumNumber
					}
				}
			}
		}
	}
	fmt.Println(targetPrice, "円を支払うためのコインの最小枚数：", minimumNumberList[coinNum-1][targetPrice])
	// 15 円を支払うためのコインの最小枚数： 2

	fmt.Println(minimumNumberList)
	// [
	// 	[0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15]  // 1
	// 	[0 1 1 2 2 3 3 4 4 5 5 6 6 7 7 8]  // 2
	// 	[0 1 1 2 2 3 3 1 2 2 3 3 4 4 2 3]  // 7
	// 	[0 1 1 2 2 3 3 1 1 2 2 3 3 4 2 2]  // 8
	// 	[0 1 1 2 2 3 3 1 1 2 2 3 1 2 2 2]  // 12
	// 	[0 1 1 2 2 3 3 1 1 2 2 3 1 2 2 2]  // 50
	// ]
}

// 動的計画法の部分で、GPTに貰った改善案
for i := 0; i < numCoins; i++ {  // コインを１枚ずつ回していく
	for j := 0; j <= targetPrice; j++ {  // priceを１つずつ回していく
		if j == 0 {
			minCoins[i][j] = 0
		} else if coinValues[i] == 1 {
			minCoins[i][j] = j
		} else if j >= coinValues[i] {  // コインの額がpriceより大きい場合は、そのコインは使用できないので、使用できる場合のみ計算処理を行う
			// 現在のコインを使用した場合
			useCoin := 1 + minCoins[i][j-coinValues[i]]
			// 使用しない場合
			noUseCoin := minCoins[i-1][j]
			// 最小枚数を選択
			if useCoin < noUseCoin {
				minCoins[i][j] = useCoin
			} else {
				minCoins[i][j] = noUseCoin
			}
		} else {
			// 現在のコインが使えない場合
			minCoins[i][j] = minCoins[i-1][j]
		}
	}