package main

import (
	"fmt"
)

type node struct {
	val int
	parent *node  // 親を持たない場合（根の場合）、nilとする
	height int
}

// 下記２点の改善を提案された
// ・経路圧縮の導入
// ・１つ関数が担当する処理をできるだけ少なくする（関数の分離）

// 互いに素な集合（１つの要素が複数の集合に属することがないデータ構造）を作成するプログラム
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
		if command[0] == 0 {  // unite（２つの要素が属する集合を合併）
			unite(command[1], command[2], nodeList)
		} else {  // same（２つの要素が同じ集合に属しているか判定する）
			if same(command[1], command[2], nodeList) {
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
		// index:  3  頂点のval:  3  親:  1  高さ:  0
		// index:  4  頂点のval:  4  親:  1  高さ:  0
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

// 引数で受け取った頂点が属する集合の代表（根）を求める
func findSet(x int, nodeList []*node) *node {
	if nodeList[x].parent == nil {  // 代表（根）の場合
		return nodeList[x]  // 代表（根）をそのまま返してあげる（再帰の基底条件）
	} else {  // 代表（根）以外の場合
		// あるノードから根までの経路上に存在する全てのノードの親に、代表（根）を設定する
		// これにより経路圧縮が実現する
		nodeList[x].parent = findSet(nodeList[x].parent.val, nodeList)  // 経路圧縮
		return nodeList[x].parent  // 常に代表を返してあげることで、呼び出した全てのノードの親に代表がセットされる
	}
}

// 同じ集合に属しているか否かを判定する
func same(x, y int, nodeList []*node) bool {
	// xとyの代表を求め、それらを比較する
	return findSet(x, nodeList) == findSet(y, nodeList)
}

// 指定された２つの要素を合併する
// unite関数による集合の合併時に高さを比較し、高さの小さいrootを持つ集合を、高さの大きいrootを持つ集合にマージするのは、
// あくまで合併後の木の高さが平坦になり、バランスの良い木を構築するために行っていることであり、
// 仮にrootの高さに不整合があったとしても（２階層の子を持ち、本来であれば高さ２だが、５の高さを持っているとされているrootが、３階層の子を持つ高さ３のrootより高さが大きいと判定されてしまう場合があったとしても）
// マージ後の木のバランスが多少悪くなるだけであり、
// しかし、これは続くfindSet関数による経路圧縮で修正される可能性があり、だからこそ何も問題がない
func unite(x, y int, nodeList []*node) {
	rootX := findSet(x, nodeList)  // xが属する集合の代表を求める
	rootY := findSet(y, nodeList)  // yが属する集合の代表を求める
	if rootX == rootY {  // 代表が同じである場合、既に同じ集合に属しているので、何もしない
		return
	}
	if rootX.height < rootY.height {
		rootX.parent = rootY
	} else if rootX.height > rootY.height {
		rootY.parent = rootX
	} else {
		rootY.parent = rootX
		rootX.height++
	}
}