package main

import (
	"fmt"
)

// --------------------
// マージソート
// --------------------
// 既にソート済みの２つの配列をマージするアルゴリズム

var compareCount int

func merge(left, right []int) []int {
	var result []int
	i, j := 0, 0
	for i < len(left) && j < len(right) {  // 片方の要素が尽きるまでまわす
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
		compareCount++
	}
	result = append(result, left[i:]...)  // leftに要素が残っている場合
	result = append(result, right[j:]...)  // rightに要素が残っている場合
	return result
}

func mergeSort(items []int) []int {
	if len(items) == 1 {  // items内の要素が１つになったら、返してあげる（再帰の基底条件）
		return items
	}
	// items内の要素が複数ある場合は、２分割して、それぞれで再帰的に処理する
	middle := len(items) / 2
	left := mergeSort(items[:middle])
	right := mergeSort(items[middle:])
	return merge(left, right)
}

func mergeSortMain() {
	// array := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	array := []int{8, 5, 9, 2, 6, 3, 7, 1, 10, 4}
	sortedArray := mergeSort(array)
	fmt.Println("Sorted array : ", sortedArray)  // Sorted array :  [1 1 2 3 3 4 5 5 5 6 9]
	fmt.Println("compareCount : ", compareCount)  // 9
}