package main

import (
	"fmt"
)

func partition() {
	array := []int{13, 19, 9, 5, 12, 8, 7, 4, 21, 2, 6, 11}
	comparingElements := array[len(array)-1]  //  ピボットの選択  全ての要素を末尾の要素である11と比較させる
	smallRangeTrailingKey := -1  // ピボットの値（11）よりも小さい値が格納された範囲を示す  0の場合、0番目の要素のみピボットの値より小さいと判断する  なので-1で初期化
	for key, value := range array {
		if value <= comparingElements {  // 各値がピボット値以下の場合
			smallRangeTrailingKey++  // ピボット値より小さい値が格納された範囲を１つ増やす
			// 新たに増えた小さい値を格納できる領域に、今回の「ピボット値より小さい」と判定された値を入れる
			// ピボット値より大きい値であるsmallRangeTrailingKeyが指す値と場所を交換してあげる
			array[smallRangeTrailingKey], array[key] = value, array[smallRangeTrailingKey]
		}
	}
	fmt.Println(array)  // [9 5 8 7 4 2 6 11 21 13 19 12]
}