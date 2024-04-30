package main

import (
	"fmt"
)

// 最長共通部分列
func longestCommonSubsequence() {
	// x := []string{"a", "b", "c", "b", "d", "a", "b"}
	// y := []string{"b", "d", "c", "a", "b", "a"}
	// 出力：4

	// x := []string{"a", "b", "c"}
	// y := []string{"a", "b", "c"}
	// 出力：３

	x := []string{"a", "b", "c"}
	y := []string{"b", "c"}
	// 出力２

	memo := make([][]int, len(x))

	for i := range memo {
		memo[i] = make([]int, len(y))
	}

	// 二重ループから分かる通り、xの長さをn、yの長さをmとすると、O(nm)のアルゴリズムである
	for xKey, xValue := range x {
		for yKey, yValue := range y {
			if xValue == yValue {  // セルの左上の値に１を足す
				if xKey == 0 || yKey == 0 {
					// 左上のセルは存在しないので、０に１を足す代わりに、１を代入する
					memo[xKey][yKey] = 1
				} else {
					// 左上のセルの値に１を足す
					memo[xKey][yKey] = memo[xKey-1][yKey-1] + 1
				}
			} else {  // 上のセル、あるいは左のセルのうち、値が大きい方のセルの値を代入する
				if xKey == 0 && yKey == 0 {
					// 上のセルも左のセルも存在しないので、０を格納する
					memo[xKey][yKey] = 0
				} else if xKey == 0 {
					// 上のセルは存在しないので、左のセルを代入する
					memo[xKey][yKey] = memo[xKey][yKey-1]
				} else if yKey == 0 {
					// 左のセルは存在しないので、上のセルを代入する
					memo[xKey][yKey] = memo[xKey-1][yKey]
				} else {
					// 上のセル、あるいは左のセルのうち、値が大きい方のセルの値を代入する
					if memo[xKey-1][yKey] >= memo[xKey][yKey-1] {
						memo[xKey][yKey] = memo[xKey-1][yKey]
					} else {
						memo[xKey][yKey] = memo[xKey][yKey-1]
					}
				}
			}
		}
	}

	// 右下のセルが、最長共通部分列の長さである
	fmt.Println(memo[len(x)-1][len(y)-1])
}