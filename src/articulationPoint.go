package main

import (
	"fmt"
)

var prenumCount int

// 関節点
func articulationPoint() {
	// n := 4  // 頂点数（０,１,２,３）
	// input := [][]int{
		// {0, 1},  // 頂点０と頂点１を結ぶ辺が存在することを意味
		// {0, 2},
		// {1, 2},
		// {2, 3},
	// }
	n := 8  // 頂点数（０,１,２,３）
	input := [][]int{
		{0, 1},
		{0, 2},
		{0, 3},
		{1, 2},
		{2, 3},
		{3, 4},
		{3, 5},
		{5, 6},
		{5, 7},
		{6, 7},
	}
	adjList := make([][]int, n)
	for i, _ := range adjList {
		adjList[i] = make([]int, 0, n)
	}
	for _, edge := range input {
		adjList[edge[0]] = append(adjList[edge[0]], edge[1])
		adjList[edge[1]] = append(adjList[edge[1]], edge[0])
	}
	// var isVisited [4]bool
	// var prenum [4]int  // 各頂点の訪問（発見）順を記録
	// var parent [4]int  // 各頂点の親頂点を記録
	var isVisited [8]bool
	var prenum [8]int
	var parent [8]int
	queue := make([]int, 0, n)
	// 頂点０を始点として深さ優先探索を開始する
	isVisited[0] = true
	parent[0] = -1
	dfs(0, adjList, &isVisited, &prenum, &parent, &queue)
	// 関節点の発見に用いるlowestというデータを求める
	// var lowest [4]int
	var lowest [8]int
	for len(queue) > 0 {
		// lowestは、訪問が完了した頂点から求めることができる
		v := queue[0]
		queue = queue[1:]
		min := prenum[v]
		var child []int
		for c, p := range parent {
			if p == v {
				child = append(child, c)
			}
		}
		backEdgeVertex := make([]int, 0, n)
		for _, dst := range adjList[v] {  // 頂点vと隣接している頂点を１つ１つdstに入れていく
			backgroundEdge := true
			for _, e := range child {  // childには、dfs treeでの子が格納されている
				if e == dst {
					backgroundEdge = false  // 該当のdstは、dfs treeで直接の親子関係であり、辺が存在する
					break
				}
			}
			if backgroundEdge {
				backEdgeVertex = append(backEdgeVertex, dst)
			}
		}
		for _, item := range backEdgeVertex {
			if min > prenum[item] {
				min = prenum[item]
			}
		}
		if len(child) > 0 {
			for _, c := range child {
				if min > lowest[c] {
					min = lowest[c]
				}
			}
		}
		lowest[v] = min
	}
	articulationPointList := make([]int, 0, n)
	// var isArticulationPoint [4]bool  // 各頂点が既に関節点として判定されているか否かを管理する
	var isArticulationPoint [8]bool
	var isRootArticulationPoint bool  // rootが関節点として判断されているか判定する用
	for i := 1; i < n; i++ {  // 頂点０はrootであり、親が存在しないので、処理の対象外
		if isArticulationPoint[parent[i]] {
			// 既に関節点と判断されているので、処理の重複を防ぐ
			continue
		} else if prenum[parent[i]] <= lowest[i] {
			articulationPointList = append(articulationPointList, parent[i])
			isArticulationPoint[parent[i]] = true
			if parent[i] == 0 {
				isRootArticulationPoint = true
			}
		}
	}
	if !isRootArticulationPoint {  // rootが上の条件により関節点と判断されたのであれば、この分岐内の処理を行う必要はない
		var rootChildCount int
		for _, p := range parent {
			if p == 0 {
				rootChildCount++
				if rootChildCount >= 2 {
					articulationPointList = append(articulationPointList, 0)
				}
			}
		}
	}
	fmt.Println(len(articulationPointList))  // 2    // 3
	fmt.Println(articulationPointList)  // [0,2]    // [0, 3, 5]
}

// func dfs(vertex int, adjList [][]int, isVisited *[4]bool, prenum *[4]int, parent *[4]int, queue *[]int) {
func dfs(vertex int, adjList [][]int, isVisited *[8]bool, prenum *[8]int, parent *[8]int, queue *[]int) {
	(*isVisited)[vertex] = true
	prenumCount++
	(*prenum)[vertex] = prenumCount
	for _, next := range adjList[vertex] {
		if !(*isVisited)[next] {
			(*parent)[next] = vertex
			dfs(next, adjList, isVisited, prenum, parent, queue)
		}
	}
	*queue = append(*queue, vertex)
}