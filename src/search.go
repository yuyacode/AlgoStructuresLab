package main

import (
	"fmt"
)

// --------------------
// 線形探索
// --------------------

// 下記コードにおける最悪の場合の時間複雑性は O(n*m)
// １つの配列内から１つのデータを線形探索で探索する場合、最悪の場合の時間複雑性は O(n) nは配列の要素数となる

func linearSearch() {
	// 下記２つの数列に共通して存在する要素を線形探索で調べる
	numberSequenceS := []int{1, 2, 3, 4, 5}
	numberSequenceQ := []int{3, 4, 1}

	count := 0
	for _, seqQVal := range numberSequenceQ {
		for _, seqSVal := range numberSequenceS {
			if seqSVal == seqQVal {
				count++
				break  // 数列内の要素に重複はないという前提なので、1回条件が成立したら、これ以上は条件が成立しない（これ以上同じ要素はない）ことになるので、breakで抜けてしまう
			}
		}
	}
	fmt.Println(count)  // 3
}


// --------------------
// 二分探索
// --------------------
func binarySearch() {
	numberSequenceS := []int{4, 5, 6, 7, 7, 9, 10, 11, 12}  // 9個の要素
	numberSequenceQ := []int{3, 5, 40, 6, 8}
	count := 0
	for _, q := range numberSequenceQ{
		low, high := 0, len(numberSequenceS) - 1
		for low <= high {  // highとlowが逆転した瞬間に、数列Qに存在する該当の要素は数列Sには存在しないと判断できる。なので、それが成立するまでは二分で探索してあげる
			middle := (low + high) / 2
			if q == numberSequenceS[middle] {
				fmt.Println("一致")
				count++
				break  // 成立したら抜けないと無限ループが発生する
			} else if q < numberSequenceS[middle] {
				fmt.Println("まだまだ小さい")
				high = middle - 1
			} else {
				fmt.Println("まだまだ大きい")
				low = middle + 1
			}
		}
	}
	fmt.Println(count)  // 2
}