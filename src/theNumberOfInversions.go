package main

import (
	"fmt"
)

// 反転数を求める

// 反転数（Inversion Count）とは、配列内の二つの要素のペアの順序が逆（降順）になっている数
// この値は、配列がどれだけ「ソートされていないか」を数値的に示す

func theNumberOfInversions() {
	// 前提：配列の要素は全て異なる値
	data := []int{5, 3, 6, 2, 1, 4}
	// data := []int{3, 5, 2, 1, 4}
	dataLength := len(data)

	// 配列をいくつかに分割して、それらの配列内での反転数を求めた後、分割した配列を統合しつつ、統合時にも反転数を求める
	// 分割した配列で算出された反転数と、統合時に算出された反転数を合計することで、元の配列の反転数を求めることができる

	// 今回は２分割にする
	left  := data[:dataLength/2]  // [5, 3, 6]
	right := data[dataLength/2:]  // [2, 1, 4]

	// 分割により得られた配列（left）において、ソートしつつ、反転数を求める
	leftReverseNumber := 0
	for i := 1; i < len(left); i++ {  // 挿入ソートを用いる
		MoveTarget := left[i]
		j := i - 1
		for j >= 0 && left[j] > MoveTarget {
			left[j+1] = left[j]
			j--
			// 反転数をカウントしていく
			// 反転数とは、要素が降順になってしまっているペアの数なので、とある要素から見たときに、その要素の左側に存在するより大きい要素の数となる
			// ここでは、それらを交換しているので、その数を数えることで反転数となる
			leftReverseNumber++
		}
		left[j+1] = MoveTarget
	}

	// 分割により得られた配列（right）において、ソートしつつ、反転数を求める
	rightReverseNumber := 0
	for i := 1; i < len(right); i++ {  // 挿入ソートを用いる
		MoveTarget := right[i]
		j := i - 1
		for j >= 0 && right[j] > MoveTarget {
			right[j+1] = right[j]
			j--
			// 反転数をカウントしていく
			// 反転数とは、要素が降順になってしまっているペアの数なので、とある要素から見たときに、その要素の左側に存在するより大きい要素の数となる
			// ここでは、それらを交換しているので、その数を数えることで反転数となる
			rightReverseNumber++
		}
		right[j+1] = MoveTarget
	}

	var sortedArray []int
	var wholeReverseNumber int
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			sortedArray = append(sortedArray, left[i])
			i++
		} else {
			sortedArray = append(sortedArray, right[j])
			j++
			// rightの要素がappendされたときに、まだappendされていないleft内の要素は、ここでappendされたrightの要素より値が大きかったということ
			// それにも関わらずleftに格納されているということは、反転が起こっていたということなので、それが起きていた要素の数をカウントする
			wholeReverseNumber += len(left) - i
		}
	}

	// leftが余っている場合、後ろにappendしてあげる
	sortedArray = append(sortedArray, left[i:]...)
	// rightが余っている場合、後ろにappendしてあげる
	sortedArray = append(sortedArray, right[j:]...)

	// 最後に、分割した配列で生じていた反転数と、統合時に特定できた反転数を合計する。この合計値が元の配列の反転数となる
	wholeReverseNumber += leftReverseNumber + rightReverseNumber

	fmt.Println(sortedArray)  // [1 2 3 4 5 6]
	fmt.Println(wholeReverseNumber)  // 10
}