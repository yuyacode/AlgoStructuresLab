package main

import (
	"fmt"
)

func fibonacciNumber(n int) int {
	// 問題文より、０項目と１項目は１だとわかっているため、計算せずとも答えを返してあげる
	if n == 0 || n == 1 {
		return 1
	}
	memo := make([]int, n+1)
	memo[0], memo[1] = 1, 1
	for i := 2; i <= n; i++ {
		// ２から上に向かって計算していくことで、常にmemoから取り出した値だけを使って答えを求めることができる
		// すなわち、求めたい解は、その直前に求めた解と、さらにその直前に求めた解の合計になる
		// つまり、O(n-2)のアルゴリズムである
		// ただし、標準的なビッグオー記法では省略されて単にO(n)と表現される
		// ビッグオー記法は主に最悪の場合の増加率を表し、具体的な計算回数の詳細を省略するためのものである
		memo[i] = memo[i-1] + memo[i-2]
	}
	return memo[n]
}

func fibonacciNumberHandler() {
	var n int
	fmt.Print("n = ")
	fmt.Scan(&n)
	result := fibonacciNumber(n)
	fmt.Println(result)
}