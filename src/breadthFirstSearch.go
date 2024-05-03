package main

import (
	"fmt"
)

type graph struct {
	vertices int     // 頂点の数
	adjList [][]int  // 隣接リスト
	distance []int   // 各頂点の始点からの距離
	queue []int      // キュー
	isDistanceCalculated []bool   // 各頂点に関して、始点からの距離を計算済みか否か
}

func NewGraph(vertices int) *graph {
	return &graph{
		vertices: vertices,
		adjList: make([][]int, vertices+1),
		distance: make([]int, vertices+1),
		queue: make([]int, 0, vertices),
		isDistanceCalculated: make([]bool, vertices+1),
	}
}

func (g *graph) createAdjList(data [][]int) {
	for _, v := range data {
		g.adjList[v[0]] = v[2:]
	}
	// fmt.Println(g.adjList)  // [[] [2 4] [4] [] [3]]
}

func (g *graph) BFS() {
	if len(g.queue) > 0 {
		visitingVertices := g.queue[0]
		g.queue = g.queue[1:]
		for _, v := range g.adjList[visitingVertices] {
			if !g.isDistanceCalculated[v] {
				g.queue = append(g.queue, v)
				g.distance[v] = g.distance[visitingVertices] + 1
				g.isDistanceCalculated[v] = true
			}
		}
		g.BFS()
	}
}

func breadthFirstSearch() {
	data := [][]int{
		// { 頂点番号, 出次数, 隣接する頂点... }
		{1, 2, 2, 4},
		{2, 1, 4},
		{3, 0},
		{4, 1, 3},
	}
	graphInstance := NewGraph(len(data))
	graphInstance.createAdjList(data)

	graphInstance.queue = append(graphInstance.queue, data[0][0])  // 始点をキューに入れる
	graphInstance.isDistanceCalculated[data[0][0]] = true  // 始点は０とする。既に０で初期化されているので計算済みとする
	graphInstance.BFS()

	fmt.Println(graphInstance.distance)  // [0 0 1 2 1]
}