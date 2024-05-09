package main

import (
	"fmt"
)

// トポロジカルソート（幅優先探索を用いて解く）
func topologicalSortBFS() {
	n := 6  // 頂点数（頂点は0からn-1）
	input := [][2]int{  // 入力データ
		{0, 1},  // 頂点０から頂点１への辺が存在することを意味する
		{1, 2},
		{3, 1},
		{3, 4},
		{4, 5},
		{5, 2},
	}
	adjList := make([][]int, n)
	for i, _ := range adjList {
		// 自己ループはないとのことなので、自身の頂点から自身の頂点宛の辺はなく、つまり容量はn-1で良い
		// 最大でn-1個の辺しか出ていかないので、
		adjList[i] = make([]int, 0, n-1)
	}
	var inDegreeInfo [6]int  // 各頂点の入次数を管理するスライス
	for _, edgeInfo := range input {
		adjList[edgeInfo[0]] = append(adjList[edgeInfo[0]], edgeInfo[1])  // 隣接リストの作成
		inDegreeInfo[edgeInfo[1]]++  // 入次数をカウントしていく
	}
	// fmt.Println(inDegreeInfo)  // [0 2 2 0 1 1]
	var queue []int
	for v, inDegree := range inDegreeInfo {
		// 入次数が０ということは、他の頂点に依存せず、はじめからトポロジカルソートを行うことができる頂点ということ
		if inDegree == 0 {
			queue = append(queue, v)
		}
	}
	// トポロジカルソートされた頂点が入ってくるスライス
	// 後々appendで足していきたいので、スライスを使う
	topologicalSortedInfo := make([]int, 0, n)
	// 入次数０の頂点から訪問するという幅優先探索を行う
	// あるいは、依存する頂点がトポロジカルソートされたことで、トポロジカルソートされる準備が整った頂点である
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		topologicalSortedInfo = append(topologicalSortedInfo, v)
		for _, j := range adjList[v] {
			// トポロジカルソートされた頂点から出ている辺の宛先頂点の入次数を-1してあげる
			// これにより、宛先頂点がトポロジカルソートされるために必要な「トポロジカルソート済みの頂点の数」が１減る
			inDegreeInfo[j]--
			if inDegreeInfo[j] == 0 {
				queue = append(queue, j)
			}
		}
	}
	fmt.Println(topologicalSortedInfo)  // [0 3 1 4 5 2]
}