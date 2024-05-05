package main

import (
	"fmt"
	"math"
)

// 単一始点最短経路
// （ダイクストラのアルゴリズム 手順）
// 初期化
// ・始点の距離を0とし、他のすべての頂点の距離を無限大（または非常に大きな数値）とする
// ・始点を開始点として設定する
// 頂点の選択と更新
// ・未確定の頂点の中から、現時点で最も短い距離にある頂点を選択
// ・選択した頂点から辿れる各頂点について、既に記録されている距離と新しい経路の距離を比較し、小さい方で更新する
// 繰り返し
// ・全ての頂点が「確定」されるまで手順2を繰り返す。頂点が確定されるとは、その頂点が最短距離で到達できると確定された状態を意味する
// 終了
// ・全ての頂点が確定された時点で、各頂点までの最短距離が求まっている

var vertexLength int
var adjMatrix [5][5]int
var shortestDistance = [5]int{math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64}
var isFixed = [5]bool{}  // 各頂点の最短経路が確定したか否かを管理する配列
var fixedVertexLength int  // 最短距離が確定した頂点の数

func execution(baseVertex int) {
	if vertexLength == fixedVertexLength {  // 頂点の数と最短距離が確定した頂点の数が一致したら、全ての頂点において始点からの最短距離が求められ、もうやることはないのでreturnする
		return
	}
	minWight := math.MaxInt64
	var nextBaseVertex int
	for v, w := range adjMatrix[baseVertex] {
		if w != 0 && !isFixed[v] {
			if shortestDistance[v] < minWight {
				minWight = shortestDistance[v]
				nextBaseVertex = v
			}
		}
	}
	isFixed[nextBaseVertex] = true
	fixedVertexLength++
	isEndPoint := true
	for v, w := range adjMatrix[nextBaseVertex] {
		if w != 0 && !isFixed[v] {
			// この関数内で最短距離が確定した、次に訪れるべき頂点nextBaseVertexに、最短距離が未確定の隣接する頂点が存在する場合は、isEndPointをfalseにする
			// そうではない場合は、nextBaseVertexの親頂点をexecution関数に渡し、自分（nextBaseVertex）以外の最短距離が確定していない頂点がないか、ある場合はそれを求め確定してもらう
			isEndPoint = false
			if shortestDistance[v] > shortestDistance[nextBaseVertex] + w {
				shortestDistance[v] = shortestDistance[nextBaseVertex] + w
			}
		}
	}
	if isEndPoint {
		execution(baseVertex)
	} else {
		execution(nextBaseVertex)
	}
}

func singleSourceShortestPath1() {
	inputSample := [][]int{
		// キー：値の意味
		// 0: 頂点番号
		// 1: 出次数
		// 2: 隣接する頂点番号
		// 3: 隣接する頂点とを繋ぐ辺の重み
		// 4: 2,3のペアが隣接する頂点の数だけ存在する
		{0, 3, 2, 3, 3, 1, 1, 2},
		{1, 2, 0, 2, 3, 4},
		{2, 3, 0, 3, 3, 1, 4, 1},
		{3, 4, 2, 1, 0, 1, 1, 4, 4, 3},
		{4, 2, 2, 1, 3, 3},
	}
	vertexLength = len(inputSample)
	for _, vertexData := range inputSample {
		vertexNum := vertexData[0]
		degree := vertexData[1]
		for i := 1; i <= degree; i++ {
			adjMatrix[vertexNum][vertexData[2*i]] = vertexData[2*i+1]
		}
	}
	fmt.Println(adjMatrix)  // 隣接行列の完成
	// [
	// 	[0 2 3 1 0] 
	// 	[2 0 0 4 0] 
	// 	[3 0 0 1 1] 
	// 	[1 4 1 0 3]
	// 	[0 0 1 3 0]
	// ]
	startVertex := 0  // 頂点０を開始点とする
	shortestDistance[startVertex] = 0  // 開始点である頂点０の最短距離を０とする
	isFixed[startVertex] = true  // 開始点である頂点０の最短距離はこれで確定した
	fixedVertexLength++
	for v, w := range adjMatrix[startVertex] {
		if w != 0 {
			shortestDistance[v] = w  // 開始点である頂点０に隣接する頂点の「始点からの距離」を現時点での最短距離として扱う
		}
	}
	execution(startVertex)
	for vertexNum, ShortestPath := range shortestDistance {
		fmt.Println(vertexNum, " ", ShortestPath)
		// 0   0
		// 1   2
		// 2   2
		// 3   1
		// 4   3
	}
}