package main

import (
	"fmt"
	"sort"
)

// 各点を表す構造体
type Point struct {
	id, x, y int
}

// k-d木の各ノードを表す構造体で、Pointのポインタ、左右の子ノードへのポインタ、そして分割を行う軸（0がx軸、1がy軸）を持つ
type KDNode struct {
	point  *Point
	left   *KDNode
	right  *KDNode
	axis   int
}

// ポイントのリストと現在の深さを受け取り、再帰的にk-d木を構築する
func buildKDTree(points []*Point, depth int) *KDNode {
	if len(points) == 0 {
		return nil
	}

	// 分割軸を決定する。深さによってx軸またはy軸を選択する
	axis := depth % 2
	// Sort points based on the current axis (0 for x, 1 for y)
	if axis == 0 {  // 深さが偶数の場合はx軸
		sort.Slice(points, func(i, j int) bool { return points[i].x < points[j].x })  // pointsスライスをx座標の昇順にソート
	} else {  // 深さが奇数の場合はy軸
		sort.Slice(points, func(i, j int) bool { return points[i].y < points[j].y })  // pointsスライスをy座標の昇順にソート
	}

	median := len(points) / 2
	node := &KDNode{
		point:  points[median],
		axis:   axis,
	}

	node.left = buildKDTree(points[:median], depth+1)
	node.right = buildKDTree(points[median+1:], depth+1)
	return node
}

func rangeSearch(root *KDNode, minX, maxX, minY, maxY int, results *[]int) {
	if root == nil {
		return
	}

	x, y := root.point.x, root.point.y
	if minX <= x && x <= maxX && minY <= y && y <= maxY {
		*results = append(*results, root.point.id)
	}

	if root.axis == 0 { // x-axis
		if minX <= x {
			rangeSearch(root.left, minX, maxX, minY, maxY, results)
		}
		if x <= maxX {
			rangeSearch(root.right, minX, maxX, minY, maxY, results)
		}
	} else { // y-axis
		if minY <= y {
			rangeSearch(root.left, minX, maxX, minY, maxY, results)
		}
		if y <= maxY {
			rangeSearch(root.right, minX, maxX, minY, maxY, results)
		}
	}
}

// 領域探索
func rangeSearchMain() {
	points := []*Point{
		{id: 0, x: 2, y: 1},
		{id: 1, x: 2, y: 2},
		{id: 2, x: 4, y: 2},
		{id: 3, x: 6, y: 2},
		{id: 4, x: 3, y: 3},
		{id: 5, x: 5, y: 4},
	}

	root := buildKDTree(points, 0)

	// Query 1
	results := []int{}
	rangeSearch(root, 2, 4, 0, 4, &results)
	fmt.Println("Query 1 results:", results) // Expect: 0, 1, 2, 4

	// Query 2
	results = []int{}
	rangeSearch(root, 4, 10, 2, 5, &results)
	fmt.Println("Query 2 results:", results) // Expect: 2, 3, 5
}
