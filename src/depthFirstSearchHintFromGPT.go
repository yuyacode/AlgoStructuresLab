package main

import (
	"fmt"
)

// グラフを表す型を作る
type Graph struct {
	vertices int     // 頂点の数を示すフィールド
	adjList [][]int  // 隣接行列を示すフィールド
	visited []bool   // 各頂点に関して、訪れたか否かを管理するフィールド
}

func NewGraph(vertices int) *Graph {
	return &Graph {
		vertices: vertices,
		adjList: make([][]int, vertices+1),  // １オリジンを想定しているため
		visited: make([]bool, vertices+1),   // １オリジンを想定しているため
	}
}

func (g *Graph) AddEdge(vertices1, vertices2 int) {
	// 隣接行列を作成する
	g.adjList[vertices1] = append(g.adjList[vertices1], vertices2)
	g.adjList[vertices2] = append(g.adjList[vertices2], vertices1)
}

func (g *Graph) DFS(vertex int) {  // 次に訪れる頂点を引数で受け取る
	g.visited[vertex] = true

	fmt.Println(vertex)  // 1, 2, 4, 3

	for _, v := range g.adjList[vertex] {
		if !g.visited[v] {  // 隣接している頂点にまだ訪れたことがない場合は、DFSメソッドを呼び深さ優先探索をする
			g.DFS(v)
		}
	}
}

func depthFirstSearchHintFromGPT() {
	graphInstancePointer := NewGraph(4)

	// グラフに対して、辺の情報を追加していく
	graphInstancePointer.AddEdge(1, 2)
	graphInstancePointer.AddEdge(1, 3)
	graphInstancePointer.AddEdge(2, 4)
	graphInstancePointer.AddEdge(3, 4)

	graphInstancePointer.DFS(1)  // 頂点１から深さ優先探索を開始していく
}
