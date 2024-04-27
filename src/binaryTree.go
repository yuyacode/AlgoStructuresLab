package main

import (
	"fmt"
)

type node struct {
	id int
	parent int
	sibling int
	degree int  // 子の数
	kind string
}

func binaryTree() {
	binaryTreeInfo := [][]int{
		{0, 1, 4},
		{1, 2, 3},
		{2, -1, -1},
		{3, -1, -1},
		{4, 5, 8},
		{5, 6, 7},
		{6, -1, -1},
		{7, -1, -1},
		{8, -1, -1},
	}

	var binaryTreeParseResult []node

	for _, nodeInfo := range binaryTreeInfo{
		for nodeInfoKey, nodeInfoVal := range nodeInfo{
			if nodeInfoKey == 0 {
				newNode := node{
					id: nodeInfoVal,  // 節点番号の設定
					kind: "internal node",  // internal nodeで初期化しておく
				}
				// 節点番号情報だけを持ったnodeをappendする
				// binaryTreeParseResult内のnode型要素を用意していかないと、後続でキーを使ったアクセスができない
				binaryTreeParseResult = append(binaryTreeParseResult, newNode)
			}
		}
	}

	for nodeKey, nodeInfo := range binaryTreeInfo{
		if nodeKey == 0 {  // 根の場合の限定処理
			binaryTreeParseResult[nodeKey].parent = -1
			binaryTreeParseResult[nodeKey].sibling = -1
			binaryTreeParseResult[nodeKey].kind = "root"
		}
		var oneSibling int
		for nodeInfoKey, nodeInfoVal := range nodeInfo{
			if nodeInfoKey == 1 {
				// 二分木は、２つの子を持つか、子を１つも持たないかの２択なので、子を表す片方の要素（１番目）のみ確認して、それが-1であれば、その要素は子を持たない要素と判定できる
				// 追記：「二分木の各ノードは最大で二つの子を持つことができる」というのが正しい理解のようなので、１つの子を持つ可能性もあるとのこと（プログラムの修正は後回しにするが、この理解はしておく）
				if nodeInfoVal == -1 {  // 子を持たない場合
					binaryTreeParseResult[nodeKey].kind = "leaf"
				} else {  // ２つの子を持つ場合
					binaryTreeParseResult[nodeInfoVal].parent = nodeKey  // 子のparentに自身を指定
					binaryTreeParseResult[nodeKey].degree = 2
					oneSibling = nodeInfoVal  // 片方の兄弟を持っておく
				}
			} else if nodeInfoKey == 2 {
				if nodeInfoVal != -1 {
					binaryTreeParseResult[nodeInfoVal].parent = nodeKey  // 子のparentに自身を指定
					binaryTreeParseResult[oneSibling].sibling = nodeInfoVal  // １番目の要素の兄弟に２番目の要素を指定
					binaryTreeParseResult[nodeInfoVal].sibling = oneSibling  // ２番目の要素の兄弟に１番目の要素を指定
				}
			}

		}
	}

	for _, nodeParseResult := range binaryTreeParseResult{
		fmt.Println("節点情報")
		fmt.Println("node: ", nodeParseResult.id)
		fmt.Println("parent: ", nodeParseResult.parent)
		fmt.Println("sibling: ", nodeParseResult.sibling)
		fmt.Println("degree: ", nodeParseResult.degree)
		fmt.Println("kind: ", nodeParseResult.kind)
		fmt.Println("ーーーーーーーーーーーーーーーーーー")
	}

	// 節点情報
	// node:  0
	// parent:  -1
	// sibling:  -1
	// degree:  2
	// kind:  root
	// ーーーーーーーーーーーーーーーーーー
	// 節点情報
	// node:  1
	// parent:  0
	// sibling:  4
	// degree:  2
	// kind:  internal node
	// ーーーーーーーーーーーーーーーーーー
	// 節点情報
	// node:  2
	// parent:  1
	// sibling:  3
	// degree:  0
	// kind:  leaf
	// ーーーーーーーーーーーーーーーーーー
	// 節点情報
	// node:  3
	// parent:  1
	// sibling:  2
	// degree:  0
	// kind:  leaf
	// ーーーーーーーーーーーーーーーーーー
	// 節点情報
	// node:  4
	// parent:  0
	// sibling:  1
	// degree:  2
	// kind:  internal node
	// ーーーーーーーーーーーーーーーーーー
	// 節点情報
	// node:  5
	// parent:  4
	// sibling:  8
	// degree:  2
	// kind:  internal node
	// ーーーーーーーーーーーーーーーーーー
	// 節点情報
	// node:  6
	// parent:  5
	// sibling:  7
	// degree:  0
	// kind:  leaf
	// ーーーーーーーーーーーーーーーーーー
	// 節点情報
	// node:  7
	// parent:  5
	// sibling:  6
	// degree:  0
	// kind:  leaf
	// ーーーーーーーーーーーーーーーーーー
	// 節点情報
	// node:  8
	// parent:  4
	// sibling:  5
	// degree:  0
	// kind:  leaf
	// ーーーーーーーーーーーーーーーーーー

}