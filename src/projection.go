package main

import (
	"fmt"
)

// 射影
// 射影とは、「２つの座標を結ぶ辺に対して、任意の点から垂線を引いたときの交点」のこと
func projection() {
	// 後続の処理で、スライスの要素を型キャストしているが、これはあまり良くないみたい
	// ただ、ここでは入力値がintで来るので、やむを得ず型キャストしている
	edgeInfo := []int{0, 0, 3, 4}  // 点１の座標、点２の座標
	baseVertex := []int{2, 5}  // 任意の点に該当

	// ２点を結ぶ辺の傾きを求める
	// var slope float64 = (float64(edgeInfo[3]) - float64(edgeInfo[1])) / (float64(edgeInfo[2]) - float64(edgeInfo[0]))  // 4/3
	slope := (float64(edgeInfo[3]) - float64(edgeInfo[1])) / (float64(edgeInfo[2]) - float64(edgeInfo[0]))  // 4/3

	// ２点を結ぶ辺の切片を求める
	// y = slope * x + b
	// edgeInfo[3] = slope * edgeInfo[2] + b
	b1 := float64(edgeInfo[3]) - (slope * float64(edgeInfo[2]))
	// これより、２点を結ぶ辺は、y = slope * x + b1

	// 垂線の傾きは、-1/slope となる
	// つまり、y = (-1/slope) * x + b となる
	// ここでも同じように、垂線となる辺の切片bを求める
	// b = y - ((-1/slope) * x)
	b2 := float64(baseVertex[1]) - ((-1/slope) * float64(baseVertex[0]))
	// これより、２点を結ぶ辺は、y = (-1/slope) * x + b2

	// 後は交点を求めるだけ
	// 交点が射影となる
	// slope * x + b1 = (-1/slope) * x + b2
	// (slope * x) - ((-1/slope) * x) = b2 - b1
	// (slope - (-1/slope)) * x = b2 - b1
	x := (b2 - b1) / (slope - (-1/slope))
	y := slope * x + b1

	fmt.Printf("%.10f, %.10f\n", x, y)  // 3.1200000000, 4.1600000000
}