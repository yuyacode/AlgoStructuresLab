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


// --------------------
// バブルソート
// --------------------
func bubbleSort() {
	arr := []int{64, 34, 25, 12, 22, 11, 9}
	n := len(arr)
    for i := 0; i < n-1; i++ {
        // swappedはこのパスで交換が行われたかどうかを追跡する
        swapped := false
        // 最後尾からすでにソートされた部分は除外して比較
        for j := 0; j < n-i-1; j++ {
            if arr[j] > arr[j+1] {
                // 隣接する要素が逆順ならば、要素を交換
                arr[j], arr[j+1] = arr[j+1], arr[j]
                swapped = true
            }
        }
        // このパスで一度も交換がなかった場合、すでに配列はソートされている
        if !swapped {
            break
        }
    }
	fmt.Println(arr)  // [11 12 22 25 34 64 90]
}


// --------------------
// 選択ソート
// --------------------
func selectionSort() {
	arr := []int{64, 25, 12, 22, 11}
	len := len(arr)

	for i := 0; i < len - 1; i++ {  // 0, 1, 2, 3
		minIndex := i  // 現在の最小
		for j := i + 1; j < len; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}

	fmt.Println(arr)  // [11 12 22 25 64]
}