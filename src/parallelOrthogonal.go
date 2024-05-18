package main

import (
	"fmt"
	"math"
)

// 辺の傾きを求める
func calculateSlope(edgeInfo []int) int {
	// 始点のx座標と終点のx座標が同じ場合、差が0になり、integer divide by zeroというエラーが出るので、その場合は傾きとして非常に大きな値を返す
	if edgeInfo[0] == edgeInfo[2] {
		return math.MaxInt64
	}
	// （終点のy座標 - 始点のy座標）/（終点のx座標 - 始点のx座標）
	return (edgeInfo[3] - edgeInfo[1]) / (edgeInfo[2] - edgeInfo[0])
}

// 並行かどうか判定する
func areParallel(edgeInfoList []int) bool {
	slope1 := calculateSlope(edgeInfoList[0:4])
	slope2 := calculateSlope(edgeInfoList[4:8])
	return slope1 == slope2
}

// 直交かどうか求める
func arePerpendicular(edgeInfoList []int) bool {
	slope1 := calculateSlope(edgeInfoList[0:4])
	slope2 := calculateSlope(edgeInfoList[4:8])
	// 一方の辺の傾きが０で、もう一方の辺の傾きが無限大の場合、直交しているが乗算結果は０になってしまうので、それに対応した条件を追加
	if (slope1 * slope2 == -1) || (slope1 == 0 && slope2 == math.MaxInt64) || (slope1 == math.MaxInt64 && slope2 == 0) {
		return true
	} else {
		return false
	}
}

// 直線の直交・並行判定
func parallelOrthogonal() {
	input := [][]int{
		{0, 0, 3, 0, 0, 2, 3, 2},  // 辺１の始点の座標、辺１の終点の座標、辺２の始点の座標、辺２の終点の座標
		{0, 0, 3, 0, 1, 1, 1, 4},
		{0, 0, 3, 0, 1, 1, 2, 2},
	}
	for _, edgeInfoList := range input {
		if areParallel(edgeInfoList) {
			fmt.Println("並行")
		} else {
			if arePerpendicular(edgeInfoList) {
				fmt.Println("直交")
			} else {
				fmt.Println("それ以外")
			}
		}
	}
	// 並行
	// 直交
	// それ以外
}