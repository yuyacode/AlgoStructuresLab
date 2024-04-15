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