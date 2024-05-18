package main

import (
	"fmt"
	"math"
)

// 反射を求める
func reflection() {
	edgeInfo := []float64{0, 0, 3, 4}  // 点１の座標、点２の座標
	slope := slope(edgeInfo)
	b1 := edgeInfo[3] - (slope * edgeInfo[2])  // edgeInfoによる辺における y=ax+b の切片bを求める

	baseVertexList := [][]float64{
		{2, 5},
		{1, 4},
		{0, 3},
	}
	for _, baseVertex := range baseVertexList {
		b2 := baseVertex[1] - ((-1/slope) * baseVertex[0])  // 垂線における y=ax+b の切片bを求める
		projectionX, projectionY := projection(slope, b1, b2)
		var reflectionX float64
		var reflectionY float64
		if baseVertex[0] == projectionX {
			reflectionX = projectionX
		} else if baseVertex[0] > projectionX {
			reflectionX = projectionX - (baseVertex[0] - projectionX)
		} else {
			reflectionX = (projectionX - baseVertex[0]) + projectionX
		}
		if baseVertex[1] == projectionY {
			reflectionY = projectionY
		} else if baseVertex[1] > projectionY {
			reflectionY = projectionY - (baseVertex[1] - projectionY)
		} else {
			reflectionY = (projectionY - baseVertex[1]) + projectionY
		}
		fmt.Printf("%.10f, %.10f\n", reflectionX, reflectionY)
	}
	// 4.2400000000, 3.3200000000
	// 3.5600000000, 2.0800000000
	// 2.8800000000, 0.8400000000
}

// 辺の傾きを求める
func slope(edgeInfo []float64) float64 {
	// 始点のx座標と終点のx座標が同じ場合、差が0になり、integer divide by zeroというエラーが出るので、その場合は傾きとして非常に大きな値を返す
	if edgeInfo[0] == edgeInfo[2] {
		return math.MaxFloat64
	}
	// （終点のy座標 - 始点のy座標）/（終点のx座標 - 始点のx座標）
	return (edgeInfo[3] - edgeInfo[1]) / (edgeInfo[2] - edgeInfo[0])
}

// 射影を求める
func projection(slope, b1, b2 float64) (float64, float64) {
	x := (b2 - b1) / (slope - (-1/slope))
	y := slope * x + b1
	// fmt.Printf("%.10f, %.10f\n", x, y)  // 3.1200000000, 4.1600000000
	return x, y
}