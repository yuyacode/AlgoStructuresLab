package main

import (
	"fmt"
	"strconv"
	"strings"
)

type node struct {
	key    int
	parent *node
	left   *node
	right  *node
}

func binarySearchTreeInsertModifyingByGPT() {
	inputData := []string{
		"insert 30",
		"insert 88",
		"insert 12",
		"insert 1",
		"insert 20",
		"find 12",
		"insert 17",
		"insert 25",
		"find 16",
		"print",
	}
	var root *node
	for _, command := range inputData {
		if strings.HasPrefix(command, "insert") {
			parts := strings.Split(command, " ")
			value, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("Failed to convert string to int:", err)
				continue
			}
			root = insert(root, nil, value)
		} else if strings.HasPrefix(command, "find") {
			parts := strings.Split(command, " ")
			value, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("Failed to convert string to int:", err)
				continue
			}
			searchResult := find(root, value)
			fmt.Println(searchResult)
		} else if command == "print" {
			fmt.Println("Preorder traversal:")
			preorderTreeWalk(root)
			fmt.Println("Inorder traversal:")
			inorderTreeWalk(root)
			fmt.Println("Postorder traversal:")
			postorderTreeWalk(root)
		}
	}
}

// n：次に比較するノード
// n：次に比較するノードの親（つまり、今比較が終わったノード）
func insert(n *node, parent *node, key int) *node {
	if n == nil {  // 次に比較するノードがなくなった時点で、新規ノードを作成してあげる
		return &node{key: key, parent: parent}
	}
	if key < n.key {
		// １つ上の n == nil 条件が成立した際に新規作成されたノード、あるいはそれらが何度か積み重なって、既に複数階層の子を持つノードが返ってくるので、それを今比較したノード（n）のleftとして設定してあげる
		n.left = insert(n.left, n, key)
	} else {
		n.right = insert(n.right, n, key)  // ifステートメントと同じように、こちらではrightとして設定してあげる
	}
	return n
}

func find(n *node, key int) string {
	if n == nil {
		return "No"
	}
	if key == n.key {
		return "Yes"
	} else if key < n.key {
		return find(n.left, key)
	} else  {
		return find(n.right, key)
	}
}

func preorderTreeWalk(n *node) {
	if n == nil {
		return
	}
	fmt.Println(n.key)
	preorderTreeWalk(n.left)
	preorderTreeWalk(n.right)
}

func inorderTreeWalk(n *node) {
	if n == nil {
		return
	}
	inorderTreeWalk(n.left)
	fmt.Println(n.key)
	inorderTreeWalk(n.right)
}

func postorderTreeWalk(n *node) {
	if n == nil {
		return
	}
	postorderTreeWalk(n.left)
	postorderTreeWalk(n.right)
	fmt.Println(n.key)
}

// Yes
// No
// Preorder traversal:
// 30
// 12
// 1
// 20
// 17
// 25
// 88
// Inorder traversal:
// 1
// 12
// 17
// 20
// 25
// 30
// 88
// Postorder traversal:
// 1
// 17
// 25
// 20
// 12
// 88
// 30