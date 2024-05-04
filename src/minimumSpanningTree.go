package main

import (
	"fmt"
)

// 最小全域木を求める方法として主に２つのアルゴリズムがある（プリムのアルゴリズムとクラスカルのアルゴリズム）
// ここでは、プリムのアルゴリズムを用いる

// プリムのアルゴリズムは、ある頂点から開始し、最小全域木に含まれる頂点を段階的に増やしていく手法
// 重みが最小の辺を選んで頂点を追加していく
// １、開始点となる任意の頂点を選ぶ
// ２、選ばれた頂点から出ている辺の中で、最小の重みを持つ辺を選び、その辺が接続する頂点を最小全域木に追加する
// ３、最小全域木に含まれる頂点から出る辺の中で、最小全域木にまだ含まれていない頂点に接続する最小の辺を選ぶ
// ４、全ての頂点が最小全域木に含まれるまで、手順3を繰り返す

// var adjacentMatrix = [][]int{  // 問題文から与えられた隣接行列（辺が存在しない場合は-1）
// 	{-1, 2, 3, 1, -1},
// 	{2, -1, -1, 4, -1},
// 	{3, -1, -1, 1, 1},
// 	{1, 4, 1, -1, 3},
// 	{-1, -1, 1, 3, -1},
// }
var adjacentMatrix = [][]int{  // 問題文から与えられた隣接行列（辺が存在しない場合は-1）
	{-1, 10, 3, -1, 18, 11, -1},
	{10, -1, 5, 1, -1, -1, -1},
	{3, 5, -1, 2, -1, 7, 5},
	{-1, 1, 2, -1, -1, -1, 2},
	{18, -1, -1, -1, -1, 1, -1},
	{11, -1, 7, -1, 1, -1, 2},
	{-1, -1, 5, 2, -1, 2, -1},
}
var vertexLength = len(adjacentMatrix)
var isInclude = make([]bool, vertexLength)  // 各頂点が最小全域木に追加されているか否かを管理するスライス
var parentVertexRelation = make([]int, vertexLength)  // キー：子の頂点、値：親の頂点
var executionCount = 1

func execution() {
	executionCount++
	minWeight := 3000  // 辺の重さの範囲は、問題より0 <= w <= 2000なので、ここでは3000で初期化しておく
	var nextIncludeVertex, parentVertexNumber int  // 次に最小全域木に追加する頂点
	for vertexNumber, include := range isInclude {  // 手順３に該当
		if include {
			for v, w := range adjacentMatrix[vertexNumber] {
				if w != -1 && !isInclude[v] {
					if w < minWeight {
						minWeight = w
						nextIncludeVertex = v
						parentVertexNumber = vertexNumber
					}
				}
			}
		}
	}
	isInclude[nextIncludeVertex] = true
	parentVertexRelation[nextIncludeVertex] = parentVertexNumber
	// execution関数実行毎に、１つの頂点が最小全域木に追加される
	// 任意の１つの頂点は、最小全域木への頂点追加処理前に既に追加されているので、残り４つの頂点を順番に追加していけば良い
	// つまり、４回execution関数を実行すれば良いので、下記の条件下でexecution関数を実行する
	if executionCount < vertexLength {
		execution()
	}
}

// 最小全域木
func minimumSpanningTree() {
	isInclude[0] = true  // 頂点０を最小全域木に追加して、開始点を頂点０に設定する
	parentVertexRelation[0] = -1  // 開始点である頂点０の親は存在しないので、-1とする
	execution()
	var edgeWeight int
	for childVertex, parentVertex := range parentVertexRelation {
		if parentVertex != -1 {
			edgeWeight += adjacentMatrix[childVertex][parentVertex]  // 最小全域木の辺の重みの総和を求める
		}
	}
	fmt.Println(edgeWeight)  // 5  // 11
}