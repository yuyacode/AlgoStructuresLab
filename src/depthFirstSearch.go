package main

import (
	"fmt"
)

// 深さ優先探索
// 「可能な限り隣接する頂点を訪問する」という戦略に基づくグラフの探索アルゴリズム

var stack []int
var adjacencyMatrices [][]int
var visitedVertices []bool
var verticesFirstVisitedInfo []int
var verticesVisitedCompleteInfo []int
var visitedOrder int

func depthFirstSearch() {
	// 隣接リスト（頂点の値, 出次数, 隣接する頂点）
	adjacencyList := [][]int{
		{1, 2, 2, 3},
		{2, 2, 3, 4},
		{3, 1, 5},
		{4, 1, 6},
		{5, 1, 6},
		{6, 0},
	}

	// 隣接行列の縦（行）を、頂点の数＋１で初期化
	adjacencyMatrices = make([][]int, len(adjacencyList)+1)

	// 訪れた頂点を管理するもの。各要素はfalseで初期化されている
	visitedVertices = make([]bool, len(adjacencyList)+1)

	// 各頂点に関して、１度目に訪れたタイミング
	verticesFirstVisitedInfo = make([]int, len(adjacencyList)+1)

	// 各頂点に関して、訪問完了となったタイミング
	verticesVisitedCompleteInfo = make([]int, len(adjacencyList)+1)

	// 隣接行列の各行を[]int型で初期化
	// 隣接リストを見る限り、出次数の最大は２なので、容量２で初期化
	for i, _ := range adjacencyMatrices {
		if i != 0 {
			adjacencyMatrices[i] = make([]int, 0, 2)
		}
	}

	// 「各頂点がどの頂点と隣接しているか」という情報を保持するスライス（adjacencyMatrices）を作成
	// 隣接行列を作っても良かったが、０が大半だと思い、メモリ効率が悪いと感じたので、今回はそれをやめる
	for _, verticesInfo := range adjacencyList {
		if len(verticesInfo) > 2 {
			adjacencyMatrices[verticesInfo[0]] = verticesInfo[2:]
		}
	}
	// fmt.Println(adjacencyMatrices)  
	// [
	// 	0 => [] 
	// 	1 => [2 3] 
	// 	2 => [3 4] 
	// 	3 => [5] 
	// 	4 => [6] 
	// 	5 => [6] 
	// 	6 => []
	// ]

	stack = append(stack, adjacencyList[0][0])  // 最初に訪れる頂点をスタックに入れる
	visitedVertices[adjacencyList[0][0]] = true
	visitedOrder++
	verticesFirstVisitedInfo[adjacencyList[0][0]] = visitedOrder

	search()

	fmt.Println(verticesFirstVisitedInfo)
	// [0 1 2 3 9 4 5]
	fmt.Println(verticesVisitedCompleteInfo)
	// [0 12 11 8 10 7 6]
}

// スタックに入れる、入れた場所に移動、移動した場所から次の移動場所を決定、次の移動場所をスタックに入れる、次の場所へ移動する
// 次の移動場所が見当たらない、今いる場所（スタックの最後の要素）をスタックから取り除く、取り除いた後のスタック内の最後の要素に移動する

func search() {
	stackLength := len(stack)
	if stackLength > 0 {
		stayingVertices := stack[stackLength-1]  // スタック内の最後の要素は、現在滞在中（今たどり着いた）頂点となる	
		nextVerticesNotFound := true
		for _, v := range adjacencyMatrices[stayingVertices] {  // 現在滞在中の頂点と繋がっている頂点を１つ１つまわす
			// if visitedVertices[v] || v == 0 {  // adjacencyMatrices[i] = make([]int, 0, 2) で初期化したので、vが０で来ることはなくなった。なので、後半の条件は不要
			if visitedVertices[v] {  // 既に訪問したことがある場合、または訪問先がない（隣接する頂点が存在しない）場合
				// ここで行う処理はないため、本来このブロックは不要だが、まあ今は学習中なので、再度見返したときに分かりやすいように残しておく
			} else {  // まだ訪問したことがない場合
				stack = append(stack, v)
				visitedVertices[v] = true
				visitedOrder++
				verticesFirstVisitedInfo[v] = visitedOrder
				nextVerticesNotFound = false
				break
			}
		}
		if nextVerticesNotFound {
			stack = stack[:stackLength-1]
			visitedOrder++
			verticesVisitedCompleteInfo[stayingVertices] = visitedOrder
			search()
		} else {
			search()
		}
	}
}