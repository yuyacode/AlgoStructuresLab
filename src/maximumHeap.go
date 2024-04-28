package main

import (
	"fmt"
)

// ヒープは、特定の大小関係によって整列されるデータ構造

// 完全二分木を表すデータ
// 最大ヒープへと変換していく
// var data = [...]int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
var data = [...]int{5, 86, 37, 12, 25, 32, 11, 7, 1, 2, 4, 19}

// 最大ヒープ
func maximumHeap() {
	// 子を持つ最後の要素から根へと順に処理していく
	// 子を持つ最後の要素は、下記で求められる
	moveTargetKey := (len(data) / 2) - 1

	for moveTargetKey >= 0 {
		CompareAndMove(moveTargetKey)
		moveTargetKey--
	}

	fmt.Println(data)
	// [16 14 10 8 7 9 3 2 4 1]
	// [86 25 37 12 5 32 11 7 1 2 4 19]
}

func CompareAndMove(moveTargetKey int) {
	// 左の子と右の子の比較
	leftChildKey := (moveTargetKey * 2) + 1
	rightChildKey := (moveTargetKey * 2) + 2

	var compareTargetKey int

	if leftChildKey < len(data) && rightChildKey < len(data) {
		// ２つの子がいる場合
		if data[leftChildKey] >= data[rightChildKey] {  // ヒープには親子間での大小関係は存在するが、兄弟間での大小関係は存在しないので、左の子と右の子が同じ値なのであれば、ここでは左の子を比較対象とする
			compareTargetKey = leftChildKey
		} else {
			compareTargetKey = rightChildKey
		}
	} else if leftChildKey < len(data) {  // 完全二分木なので、右の子のみいる状態はあり得ない。つまり左の子のみ存在しているケースだけを考えれば良い
		// 左の子だけいる場合
		compareTargetKey = leftChildKey
	} else {
		// 子がいない場合
		return
	}

	if data[moveTargetKey] < data[compareTargetKey] {  // 最大ヒープ（親が子より大きい状態）にしたいので、既にそのようになっているものに関しては何もする必要がない。つまり。親が子より小さい場合だけ、それを逆にしてあげる
		data[moveTargetKey], data[compareTargetKey] = data[compareTargetKey], data[moveTargetKey]
		CompareAndMove(compareTargetKey)  // ここで再帰的に呼ぶ  compareTargetKeyを渡す  compareTargetKeyが次のmoveTargetKeyとして機能する
	}
	return
}