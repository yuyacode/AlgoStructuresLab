package main

import (
	"fmt"
	"strconv"
	"strings"
)

type node struct {
	key int
	parent *node
	left *node
	right *node
}

// 二分探索木
// 二分探索木とは、二分木の概念に「データの検索効率を向上させるための条件（左の子 < 親ノード < 右の子）」を加えた特殊な形の二分木。二分木の一種である。
func binarySearchTreeInsert() {
	inputData := []string{
		"insert 30",
		"insert 88",
		"insert 12",
		"insert 1",
		"insert 20",
		"insert 17",
		"insert 25",
		"print",
	}
	var root *node
	for key, command := range inputData {
		if strings.HasPrefix(command, "insert") {
			parts := strings.Split(command, " ")
			if len(parts) == 2 {
				value, err := strconv.Atoi(parts[1])
				if err != nil {
					fmt.Println("string型からint型への変換に失敗しました。", err)
				} else {
					// 二分探索木への挿入処理
					if key == 0 {
						root = &node{
							key: value,
						}
					} else {
						var parentNode *node
						if value < root.key {
							nextNode := root.left
							if nextNode != nil {
								for {
									if value < nextNode.key {
										parentNode = nextNode
										nextNode = nextNode.left
										if nextNode == nil {
											newNode := &node{
												key: value,
												parent: parentNode,
											}
											parentNode.left = newNode
											break
										}
									} else {
										parentNode = nextNode
										nextNode = nextNode.right
										if nextNode == nil {
											newNode := &node{
												key: value,
												parent: parentNode,
											}
											parentNode.right = newNode
											break
										}
									}
								}
							} else {  // root.leftがnilの場合は、そこに新しいnodeを挿入する
								newNode := &node{
									key: value,
									parent: root,
								}
								root.left = newNode
							}
						} else {
							nextNode := root.right
							if nextNode != nil {
								for {
									if value < nextNode.key {
										parentNode = nextNode
										nextNode = nextNode.left
										if nextNode == nil {
											newNode := &node{
												key: value,
												parent: parentNode,
											}
											parentNode.left = newNode
											break
										}
									} else {
										parentNode = nextNode
										nextNode = nextNode.right
										if nextNode == nil {
											newNode := &node{
												key: value,
												parent: parentNode,
											}
											parentNode.right = newNode
											break
										}
									}
								}
							} else {  // root.leftがnilの場合は、そこに新しいnodeを挿入する
								newNode := &node{
									key: value,
									parent: root,
								}
								root.right = newNode
							}
						}
					}
				}
			}
		} else if command == "print" {
			// 二分探索木の表示
			// 先行順巡回アルゴリズムによる表示
			fmt.Println("先行順巡回アルゴリズム")
			preorderTreeWalk(root)
			// 中間順巡回アルゴリズムによる表示
			fmt.Println("中間順巡回アルゴリズム")
			inorderTreeWalk(root)
			// 後行順巡回アルゴリズムによる表示
			fmt.Println("後行順巡回アルゴリズム")
			postorderTreeWalk(root)
		}
	}
}

// 先行順巡回アルゴリズム
func preorderTreeWalk(node *node) {
	if node == nil {
		return
	}
	fmt.Println(node.key)
	preorderTreeWalk(node.left)
	preorderTreeWalk(node.right)
}

// 中間順巡回アルゴリズム
func inorderTreeWalk(node *node) {
	if node == nil {
		return
	}
	inorderTreeWalk(node.left)
	fmt.Println(node.key)
	inorderTreeWalk(node.right)
}

// 後行順巡回アルゴリズム
func postorderTreeWalk(node *node) {
	if node == nil {
		return
	}
	postorderTreeWalk(node.left)
	postorderTreeWalk(node.right)
	fmt.Println(node.key)
}

// 先行順巡回アルゴリズム
// 30
// 12
// 1
// 20
// 17
// 25
// 88
// 中間順巡回アルゴリズム
// 1
// 12
// 17
// 20
// 25
// 30
// 88
// 後行順巡回アルゴリズム
// 1
// 17
// 25
// 20
// 12
// 88
// 30