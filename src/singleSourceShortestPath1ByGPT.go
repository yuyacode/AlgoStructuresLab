package main

import (
	"container/heap"
	"fmt"
	"math"
)

// src/singleSourceShortestPath1.goのコードをGPTに添削してもらい、その後に貰った改善コード
// ・最小の距離を持つ未確定の頂点を探すために全ての頂点を走査しており非効率なので、優先度キューを使用すると良いとのこと
// ・関数間で状態を共有するためにグローバル変数を使用するのは、やはり一般的に良いプラクティスではないらしく、引数や戻り値で対応しろとのこと
// ・隣接行列が適切に初期化されていない場合や、入力データが予期しないフォーマットである場合に対応するエラーハンドリングを追加すると良いとのこと

type PriorityQueue []*Vertex  // 頂点が格納されたキュー

type Vertex struct {  // 頂点自体を示す構造体
	id       int  // 頂点番号
	distance int  // 最短距離
	index    int // ヒープ内のインデックス
}

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].distance < pq[j].distance
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	vertex := x.(*Vertex)
	vertex.index = n
	*pq = append(*pq, vertex)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	vertex := old[n-1]
	vertex.index = -1
	*pq = old[0 : n-1]
	return vertex
}

// ダイクストラのアルゴリズムを実行する関数
func dijkstra(adjMatrix [][]int, startVertex int) []int {
	vertexLength := len(adjMatrix)
	distances := make([]int, vertexLength)
	for i := range distances {
		distances[i] = math.MaxInt64
	}
	distances[startVertex] = 0

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Vertex{id: startVertex, distance: 0})  // PriorityQueue型に紐づくPushメソッドが実行される（heapパッケージの挙動）

	for pq.Len() > 0 {
		u := heap.Pop(&pq).(*Vertex)  // heap.Pop(&pq)の戻り値を型アサーションしている
		for v, weight := range adjMatrix[u.id] {
			if weight > 0 && distances[v] > distances[u.id]+weight {
				distances[v] = distances[u.id] + weight
				heap.Push(&pq, &Vertex{id: v, distance: distances[v]})
			}
		}
	}

	return distances
}

func main() {
	adjMatrix := [][]int{  // 各頂点間の辺の重みを示した隣接行列
		{0, 2, 3, 1, 0},
		{2, 0, 0, 4, 0},
		{3, 0, 0, 1, 1},
		{1, 4, 1, 0, 3},
		{0, 0, 1, 3, 0},
	}
	distances := dijkstra(adjMatrix, 0)
	for i, d := range distances {
		fmt.Printf("Vertex %d, Distance: %d\n", i, d)
	}
}
