package main

import (
	"fmt"
)

var startVertex int

type graph struct {
	adjList [][]int  // 隣接リストを示すフィールド
	visited []bool   // 各頂点に関して、訪れたか否かを管理するフィールド
	searchId []int   // 各頂点に関して、どの探索により辿り着いたか（どの頂点を始点とした探索で辿り着いたか）
}

func newGraph(vertexNumber int) *graph {
	return &graph{
		adjList: make([][]int, vertexNumber),
		visited: make([]bool, vertexNumber),
		searchId: make([]int, vertexNumber),
	}
}

func (g *graph) createAdjList(vertexRelationship [][]int) {
	for _, v := range vertexRelationship {
		g.adjList[v[0]] = append(g.adjList[v[0]], v[1])
		g.adjList[v[1]] = append(g.adjList[v[1]], v[0])
	}
	// fmt.Println(g.adjList)
	// [
	// 	[1 2] 
	// 	[0] 
	// 	[0] 
	// 	[4] 
	// 	[3] 
	// 	[7 6]     
	// 	[5 7 8] 
	// 	[5 6 8] 
	// 	[6 7 9] 
	// 	[8]
	// ]
}

func (g *graph) dFS(vertex int) {
	g.visited[vertex] = true
	g.searchId[vertex] = startVertex
	for _, v := range g.adjList[vertex] {
		if !g.visited[v] {
			g.dFS(v)
		}
	}
}

// 連結成分
func connectedComponents() {
	vertexNumber := 10
	vertexRelationship := [][]int{
		{0, 1},
		{0, 2},
		{3, 4},
		{5, 7},
		{5, 6},
		{6, 7},
		{6, 8},
		{7, 8},
		{8, 9},
	}
	survey_target := [][]int{
		{0, 1},
		{5, 9},
		{1, 3},
	}
	g := newGraph(vertexNumber)
	g.createAdjList(vertexRelationship)
	for i := 0; i < 10; i++ {  // 各頂点を始点として深さ優先探索をしていく
		if !g.visited[i] {  // 頂点iに対して、既に他の頂点を始点とした探索で辿り着いている場合は、頂点iを始点とした探索を行う必要はないため、スキップする
			startVertex = i  // 始点の設定
			g.dFS(i)
		}
	}
	// fmt.Println(g.searchId)  // [0 0 0 3 3 5 5 5 5 5]
	for _, v := range survey_target {
		if g.searchId[v[0]] == g.searchId[v[1]] {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}
}