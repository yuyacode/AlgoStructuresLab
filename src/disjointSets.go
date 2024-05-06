package main

import (
	"fmt"
)

type node struct {
	val int
	parent *node  // 親を持たない場合（根の場合）、nilとする
	height int
}

// 互いに素な集合（１つの要素が複数の集合に属することがないデータ構造ぞう）
func disjointSets() {
	// 問題文より、nodeLength（頂点の数）は５と与えられている
	// ０からnodeLength-1がnodeらしいので、ここでは０から４のノードが存在することになる
	nodeLength := 5
	var nodeList = make([]*node, nodeLength)
	for i := 0; i < nodeLength; i++ {
		nodeList[i] = makeSet(i)
	}
	commandList := [][]int{
		// 0番目: 0,unite、1,same
		// 1番目: 要素1
		// 2番目: 要素2
		{0, 1, 4},
		{0, 2, 3},
		{1, 1, 2},
		{1, 3, 4},
		{1, 1, 4},
		{1, 3, 2},
		{0, 1, 3},
		{1, 2, 4},
		{1, 3, 0},
		{0, 0, 4},
		{1, 0, 2},
		{1, 3, 0},
	}
	for _, command := range commandList {
		if command[0] == 0 {  // unite
			unite(command[1], command[2], nodeList)
		} else {  // same
			findSet(command[1], command[2], nodeList)
		}
	}
	for i, n := range nodeList {
		parent := -1
		if n.parent != nil {
			parent = n.parent.val
		}
		fmt.Println("index: ", i, " 頂点のval: ", n.val, " 親: ", parent, " 高さ: ", n.height)
		// index:  0  頂点のval:  0  親:  1  高さ:  0
		// index:  1  頂点のval:  1  親:  -1  高さ:  2
		// index:  2  頂点のval:  2  親:  1  高さ:  1
		// index:  4  頂点のval:  4  親:  1  高さ:  0
		// index:  3  頂点のval:  3  親:  2  高さ:  0
	}
}

// 要素xがただ１つの集合を作る
func makeSet(val int) *node {
	return &node{
		val: val,
		parent: nil,
		height: 0,
	}
}

// 要素xが属する集合の代表要素を求める
// 代表要素が同じであれば、２つの要素は同じ集合に属すると判断できる
// 反対に、異なる場合は、異なる集合に属すると判断できる
// 指定された２つの要素が同じ集合に属するかどうかを調べる操作を Union Find という
func findSet(x, y int, nodeList []*node) {
	// 要素xが属する木の代表（root）の値を返す
	for nodeList[x].parent != nil {
		x = nodeList[x].parent.val
	}
	for nodeList[y].parent != nil {
		y = nodeList[y].parent.val
	}
	if nodeList[x].val == nodeList[y].val {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
	// 0
	// 0
	// 1
	// 1
	// 1
	// 0
	// 1
	// 1
}

// 指定された２つの要素を合併する
func unite(x, y int, nodeList []*node) {
	// ２つの要素それぞれの代表のうち、一方の代表を新しい代表として選び、代表にならなかった方の代表が新しい代表を親として指すようにする
	// 高さの低い代表を、高さの高い代表にマージする。これにより、合併後の木の高さが高くなることはない
	// 同じ高さの木を合併するときは、選ばれた代表の高さを＋１する
	// var routeCompressionTarget []*node  // 経路圧縮を行うnodeを管理するスライス  今回の実装では経路圧縮は行わない
	for nodeList[x].parent != nil {
		// ある要素の代表を求めるとき、その要素から代表に至る経路上の全てのnodeについて、ポインタが直接代表を指すように変更することで、経路圧縮を行う
		// ある要素から代表に至るまでの経路上にあるnodeが圧縮対象なので、それらの兄弟要素は圧縮対象外である
		// routeCompressionTarget = append(routeCompressionTarget, nodeList[x])
		x = nodeList[x].parent.val
	}
	for nodeList[y].parent != nil {
		// routeCompressionTarget = append(routeCompressionTarget, nodeList[y])
		y = nodeList[y].parent.val
	}
	if nodeList[x].height < nodeList[y].height {  // 求めた代表同士の比較
		nodeList[x].parent = nodeList[y]
		// for _, v := range routeCompressionTarget {
		// 	v.parent = nodeList[y]
		// 	v.height = 0
		// }
	} else if nodeList[x].height > nodeList[y].height {
		nodeList[y].parent = nodeList[x]
		// for _, v := range routeCompressionTarget {
		// 	v.parent = nodeList[x]
		// 	v.height = 0
		// }
	} else {  // ２つの代表の高さが同じだったら、今回はxを親とする
		nodeList[y].parent = nodeList[x]
		nodeList[x].height++
		// for _, v := range routeCompressionTarget {
		// 	v.parent = nodeList[x]
		// 	v.height = 0
		// }
	}
}