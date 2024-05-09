package main

import (
	"fmt"
)

// トポロジカルソート（深さ優先探索を用いて解く）
func topologicalSortDFS() {
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
	for _, edgeInfo := range input {
		adjList[edgeInfo[0]] = append(adjList[edgeInfo[0]], edgeInfo[1])  // 隣接リストの作成
	}
	isVisited := make([]bool, n)  // 各頂点に関して、訪問が完了したか否か管理するスライス
	topologicalSortedInfo := make([]int, 0, n)
	for i := 0; i < n; i++ {
		if !isVisited[i] {
			// この関数内のtopologicalSortedInfoは長さ０という情報を保持しているので、dfs関数内でappendが起こり、長さに変更が加わっても、それを検知することができない
			// つまり長さ０という認識のまま下記の出力へと移るので、dfs関数内でappendした要素が見えず、出力が起こらない
			// なので、アドレス演算子によりアドレスを渡すことで、データ本体へのポインタ、長さ情報、容量情報全てを参照渡しさせる
			// これにより、appendごとに長さの変更が加わっても、それをこの関数内から検知でき、出力が正常に行われる
			dfs(i, adjList, isVisited, &topologicalSortedInfo)
		}
	}
	// 深さ優先探索の場合、トポロジカル順が逆順に決定していくので、逆順で出力してあげる
	for i := n-1; i >= 0; i-- {
		fmt.Println(topologicalSortedInfo[i])  // 3 4 5 0 1 2
	}
	// トポロジカルソートの問題では、１つの入力に対して複数の解答が存在するケースがある
	// 今回もそのケースであり、同じ入力に対しても、幅優先探索を用いたときと結果が異なる
}

func dfs(i int, adjList [][]int, isVisited []bool, topologicalSortedInfo *[]int) {
	isVisited[i] = true
	for _, v := range adjList[i] {
		if !isVisited[v] {
			dfs(v, adjList, isVisited, topologicalSortedInfo)
		}
	}
	*topologicalSortedInfo = append(*topologicalSortedInfo, i)
}

// 下記は別件
// スライスの落とし穴について
func test() {
	topologicalSortedInfo := make([]int, 6)
	fmt.Println(topologicalSortedInfo)  // [0 0 0 0 0 0]

	test2(topologicalSortedInfo)  // topologicalSortedInfo自体のデータへのポインタ、長さ情報（６）、容量情報（６）が値渡しされる

	// test2関数内でスライスの拡張が起こり、この関数内のtopologicalSortedInfoと指しているメモリが異なるため、test2関数の変更は反映されていない
	fmt.Println(topologicalSortedInfo)  // [0 0 0 0 0 0]
}

func test2(topologicalSortedInfo []int) {
	// 容量６全て埋まっているため、スライスの拡張が起こる
	topologicalSortedInfo = append(topologicalSortedInfo, 1)
	fmt.Println(topologicalSortedInfo)  // [0 0 0 0 0 0 1]
}