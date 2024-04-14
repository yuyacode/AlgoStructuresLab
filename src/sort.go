package main

import (
	"fmt"
)

// --------------------
// 挿入ソート
// --------------------

// 挿入ソートは、各イテレーションで未ソートの領域から要素を取り出し、それをソート済みの適切な位置に挿入することで動作する

// アルゴリズムは以下の手順に従う
// 1. 先頭の要素をソート済みとする
// 2. 次の要素を取り、ソート済みの部分において正しい位置に挿入する
// 3. ソート済みのリストのサイズが全体のサイズに等しくなるまで繰り返す

func insertionSort() {
	arr := []int{10, 9, 8, 7, 6, 5}
	for i := 1; i < len(arr); i++ {
		currentMoveTarget :=  arr[i]
		j := i - 1
		for j >= 0 && arr[j] > currentMoveTarget {
			arr[j + 1] = arr[j]  // １つ後ろへずらす
			j--
		}
		arr[j + 1] = currentMoveTarget
	}
	fmt.Println(arr)  // [5 6 7 8 9 10]
}