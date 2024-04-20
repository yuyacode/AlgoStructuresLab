package main

import (
	"fmt"
)

func merge(left []int, pivotValue int, right []int) []int {
	// 空間複雑度を下げるために、leftに対して、pivotとrightをマージしてあげる
	// 別で変数を用意して、そこにleft, pivotValue, rightを統合する形には、ここではしない
	left = append(left, pivotValue)
	left = append(left, right...)
	return left
}

func quickSort(array []int) []int {
	// 再帰的に呼ぶ場合、どこかで必ずreturnさせる終点が必要（再帰の基底条件）
	// 引数で受け取った配列の要素数が１になったら、返してあげる
	// = 1 にした場合、arrayが空の時に分岐が成立せず、後続の処理へと移る。このとき、pivotValueへの代入において、インデックス-1にアクセスしてパニックを起こすので、必ず2以下という条件にしてあげる
	if len(array) < 2 {
		return array
	}
	// 引数で受け取った配列内の要素が複数ある場合は、パーティーションにより分割しつつソートする
	pivotKey := len(array)-1
	pivotValue := array[pivotKey]
	smallRangeTrailingKey := -1  // ピボットの値よりも小さい値が格納された範囲を示す  0の場合、0番目の要素のみピボットの値より小さいと判断する  なので-1で初期化
	for key, value := range array {
		if value <= pivotValue {  // 各値がピボット値以下の場合
			smallRangeTrailingKey++  // ピボット値より小さい値が格納された範囲を１つ増やす
			// 新たに増えた小さい値を格納できる領域に、今回の「ピボット値より小さい」と判定された値を入れる
			// ピボット値より大きい値であるsmallRangeTrailingKeyが指す値と場所を交換してあげる
			array[smallRangeTrailingKey], array[key] = value, array[smallRangeTrailingKey]
		}
	}
	// パーティーション後のarrayの要素数が２の場合は、既に２つの要素間でソートが完了しているので、ここで返してあげる
	// ３つ以上の場合は、まだそれら要素間でのソートが完了していないので、引き続き再帰的に呼ぶことで、ソートする
	if len(array) == 2 {
		return array
	} else {
		// 再帰的に呼ぶ
		// この時点で、array[smallRangeTrailingKey]にはピボット値が格納されている
		left := quickSort(array[:smallRangeTrailingKey])  // ピボットより左側を渡す
		right := quickSort(array[smallRangeTrailingKey+1:]) // ピボットより右側を渡す
		// 最後はマージする
		return merge(left, pivotValue, right)
	}
}

func quickSortMain() {
	array := []int{13, 19, 9, 5, 12, 8, 7, 4, 21, 2, 5, 3, 14, 6, 11}
	// array := []int{13, 19, 9, 5, 12}
	sortedArray := quickSort(array)
	fmt.Println("sortedArray : ", sortedArray)  //  [2 3 4 5 5 6 7 8 9 11 12 13 14 19 21]
}