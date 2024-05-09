package main

import (
	"fmt"
)

// 無限大を表す定数
const inf = 1<<60

func floydWarshall(dist [][]int, n int) {
	for k := 0; k < n; k++ {  // 経由点kを固定化した上で、始点iと終点jを回し、全頂点ペアにおいて経由点kを経由したときのコストを求める
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dist[i][k] < inf && dist[k][j] < inf {  // オーバーフローを避ける
					newDist := dist[i][k] + dist[k][j]
					if newDist < dist[i][j] {
						dist[i][j] = newDist
					}
				}
			}
		}
	}
}

// 全点対間最短経路
func allPairsShortestPath() {
	n := 4  // 頂点数

	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			if i == j {
				dist[i][j] = 0
			} else {
				dist[i][j] = inf
			}
		}
	}

	// input := [][]int{
	// 	{0, 1, 1},  // 頂点０から頂点１までの辺の重みが１という意
	// 	{0, 2, 5},
	// 	{1, 2, 2},
	// 	{1, 3, 4},
	// 	{2, 3, 1},
	// 	{3, 2, 7},
	// }

	// input := [][]int{
	// 	{0, 1, 1},
	// 	{0, 2, -5},
	// 	{1, 2, 2},
	// 	{1, 3, 4},
	// 	{2, 3, 1},
	// 	{3, 2, 7},
	// }

	input := [][]int{
		{0, 1, 1},
		{0, 2, 5},
		{1, 2, 2},
		{1, 3, 4},
		{2, 3, 1},
		{3, 2, -7},
	}
	for _, edgeInfo := range input {
		dist[edgeInfo[0]][edgeInfo[1]] = edgeInfo[2]
	}

	// ワーシャルフロイドアルゴリズムの実行
	floydWarshall(dist, n)

	// グラフが負の閉路を持つ場合は、NEGATIVE CYCLEと出力する
	isNegativeCycle := false
	for v := 0; v < n; v++ {
		if dist[v][v] < 0 {
			fmt.Println("NEGATIVE CYCLE")
			isNegativeCycle = true
		}
	}

	if !isNegativeCycle {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dist[i][j] == inf {  // infということは、iからjへの経路が存在しないことを指す
					fmt.Print("inf")
				} else {
					fmt.Print(dist[i][j])
				}
				fmt.Print(" ")
			}
			fmt.Print("\n")
		}
	}
	// 0 1 3 4 
	// inf 0 2 3 
	// inf inf 0 1 
	// inf inf 7 0

	// 0 1 -5 -4 
	// inf 0 2 3 
	// inf inf 0 1 
	// inf inf 7 0 

	// NEGATIVE CYCLE
	// NEGATIVE CYCLE
}