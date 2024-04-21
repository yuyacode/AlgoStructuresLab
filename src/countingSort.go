package main

import (
	"fmt"
)

// 計数ソート
func countingSort() {
	data := []int{4, 5, 0, 3, 1, 5, 0, 5}

	// 配列内に負の要素しか存在しない場合、0で初期化をすると、それが最大値となってしまう罠がある
	// 今回は負の要素が含まれないため、var max int による初期化でも良いが、癖付けるために0番目の要素で初期化する
	// 計数ソートは非負の整数だけが格納された配列を扱う
	// なぜなら、計数ソートは処理の中で、配列の要素をキーとして使用するため
	// var max int
	max := data[0]
	for _, item := range data {
		if item > max {
			max = item  // 最大値の特定
		}
	}

	countArr := make([]int, max+1)
	for _, item := range data {
		countArr[item]++  // どの値が何個存在するか確認
	}

	totalCountArr := make([]int, max+1)
	for i := 0; i <= max; i++ {
		if i == 0 {
			totalCountArr[i] = countArr[i]
		} else {
			totalCountArr[i] = totalCountArr[i-1] + countArr[i]  // 直前までの累計とその要素の合計を足すことで、累計和を算出
		}
	}

	fmt.Println(countArr)       // [2 1 0 1 1 3]
	fmt.Println(totalCountArr)  // [2 3 3 4 5 8]

	dataLength := len(data)  // 8
	sortedData := make([]int, dataLength)
	for i := dataLength-1; i >= 0; i-- {
		key := totalCountArr[data[i]]
		sortedData[key-1] = data[i]
		totalCountArr[data[i]]--
	}
	fmt.Println(sortedData)  // [0 0 1 3 4 5 5 5]
}