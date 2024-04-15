package main

import (
	"fmt"
)

// --------------------
// 水溜まりの体積を求める練習問題
// --------------------
func areaCalculation() {
	areaInfo := []string{"\\", "\\", "/", "/", "/", "\\", "_", "/"}
	stack := []int{}
	totalArea := 0
	for key, data := range areaInfo {
		if data == "\\" {
			stack = append(stack, key + 1)  // 水溜まりが発生する位置をスタックに記録
		} else if data == "/" {
			stackLength := len(stack)
			if stackLength == 0 {
				// スタックに値が存在しない場合、イメージとしては丘陵が上がるだけであり、次の水溜まりが発生する標高が高くなるだけ
				// なので、何もしない
			} else {
				puddleOccurrencePosition := stack[stackLength - 1 : ]  // 水溜まりが発生した位置を、スタックから取得
				stack = stack[ : stackLength - 1]  // ポップした分をスタックから削除
				totalArea += (key + 1) - puddleOccurrencePosition[0]
			}
		} else {
			// _の場合、何もしない
		}
	}
	fmt.Println(totalArea)  // 6
}