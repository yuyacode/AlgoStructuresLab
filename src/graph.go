package main

import (
	"fmt"
)

func graph() {
	// 頂点番号, 出次数, 隣接する頂点の番号
	dataList := [][]int{
		{1, 2, 2, 4},
		{2, 1, 4},
		{3, 0},
		{4, 1, 3},
	}

	// 隣接リスト表現の形式で与えられた有向グラフの隣接行列を出力するプログラム

	adjacencyMatrices := make([][]int, len(dataList)+1)  // インデックスが１〜４のスライスを作成したい。そのため容量５で初期化し、後続で1〜４のみ使用することにする。０は使用しないインデックスとする。

	for i, _ := range adjacencyMatrices {
		if i != 0 {
			adjacencyMatrices[i] = make([]int, len(dataList)+1)
		}
	}

	for _, data := range dataList {
		var targetNode int
		for i, v := range data {
			if i == 0 {
				targetNode = v
			} else if i >= 2 {
				adjacencyMatrices[targetNode][v] = 1
			}
		}
	}

	for lineCount, adjacencyMatricesData := range adjacencyMatrices {
		if lineCount != 0 {
			for i, sideExistence := range adjacencyMatricesData {
				if i != 0 {
					fmt.Println(sideExistence, " ")
				}
			}
		}
	}
}

// 0  1  0  1  
// 0  0  0  1  
// 0  0  0  0  
// 0  0  1  0