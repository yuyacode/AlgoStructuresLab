package main

import (
	"fmt"
	"sort"
)

type Edge struct {
	Weight, Src, Dest int
}

// グラフの頂点がどのようにグルーピングされているかを追跡する
// 初期状態では、各頂点は自分自身のみを含む独立した集合に属している
type UnionFind struct {
	parent, rank []int
}

func NewUnionFind(size int) *UnionFind {
	parent := make([]int, size)  // 各頂点の親頂点を管理
	rank := make([]int, size)  // 各集合の木構造の深さを近似的に保持する
	for i := range parent {
		parent[i] = i  // 初期状態では自身が親（すなわち自己参照）
	}
	return &UnionFind{parent: parent, rank: rank}
}

// 頂点iが属する集合の代表（root）を返す
func (uf *UnionFind) Find(i int) int {
	if uf.parent[i] != i {
		// 返ってきた値（代表）を親に設定することで、経路圧縮を実現
		uf.parent[i] = uf.Find(uf.parent[i])
	}
	// 自己参照を見つけたときに、その頂点が返る（再帰の基底条件）
	// 自己参照している頂点が、その頂点が属する集合の代表である
	return uf.parent[i]
}

// 集合を統合
func (uf *UnionFind) Union(x, y int) {
	rootX := uf.Find(x)
	rootY := uf.Find(y)

	// 最小全域木に追加する辺の両端の頂点が同じ集合に属するか確認する
	// 異なる集合に属する場合のみ、その辺を最小全域木に追加することで、グラフ内に閉路ができることを防ぐ
	if rootX != rootY {
		if uf.rank[rootX] > uf.rank[rootY] {
			uf.parent[rootY] = rootX
		} else if uf.rank[rootX] < uf.rank[rootY] {
			uf.parent[rootX] = rootY
		} else {
			uf.parent[rootY] = rootX
			uf.rank[rootX]++
		}
	}
}

func Kruskal(edges []Edge, numVertices int) int {
	sort.Slice(edges, func(i, j int) bool {  // 第二引数ではソートの順序を決定するための比較関数を提供
		return edges[i].Weight < edges[j].Weight  // この条件を渡すことで、weightの昇順でソートしている
	})
	uf := NewUnionFind(numVertices)

	result := 0
	for _, e := range edges {  // edgesが昇順で並んでいることで、重みの小さい辺から判定を行える
		// 辺の両端の頂点が属する集合の代表を確認する
		// 代表が異なる場合、それら２つの頂点は異なる集合に属していることになるので、連結しても閉路を形成せず、最小全域木の契約を守ることができる
		if uf.Find(e.Src) != uf.Find(e.Dest) {
			result += e.Weight
			uf.Union(e.Src, e.Dest)
		}
	}

	return result
}

func minimumSpanningTreeKruskal() {
	edges := []Edge{
		// {Weight: 10, Src: 0, Dest: 1},
		// {Weight: 6, Src: 0, Dest: 2},
		// {Weight: 5, Src: 0, Dest: 3},
		// {Weight: 15, Src: 1, Dest: 3},
		// {Weight: 4, Src: 2, Dest: 3},

		{Weight: 1, Src: 0, Dest: 1},
		{Weight: 3, Src: 0, Dest: 2},
		{Weight: 1, Src: 1, Dest: 2},
		{Weight: 7, Src: 1, Dest: 3},
		{Weight: 1, Src: 2, Dest: 4},
		{Weight: 3, Src: 1, Dest: 4},
		{Weight: 1, Src: 3, Dest: 4},
		{Weight: 1, Src: 3, Dest: 5},
		{Weight: 6, Src: 4, Dest: 5},
	}
	// numVertices := 4
	numVertices := 6
	minCost := Kruskal(edges, numVertices)
	fmt.Printf("The total weight of the minimum spanning tree is: %d\n", minCost)  // 5
}
