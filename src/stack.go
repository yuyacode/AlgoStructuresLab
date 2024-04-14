package main

import (
	"fmt"
)

func stack() {
	dataList := []interface{}{1, 2, "+", 3, 4, "+", "*"}  // 逆ポーランド記法をスライスで無理矢理表現
	stack := []int{}
	for _, data := range dataList {
		switch data := data.(type) {
		case int:
			stack = append(stack, data)
		case string:
			stackLength := len(stack)
			if stackLength < 2 {
				fmt.Println("演算対象が存在しません")
			} else {
				lastTwo := stack[stackLength-2:]
				stack = stack[:stackLength-2]  // ポップされた分をスタックから削除
				if data == "+" {
					stack = append(stack, lastTwo[0] + lastTwo[1])
				} else if data == "*" {
					stack = append(stack, lastTwo[0] * lastTwo[1])
				} else if data == "-" {
					stack = append(stack, lastTwo[0] - lastTwo[1])
				} else {
					fmt.Println("想定しない演算子です")
				}
			}
		default:
			fmt.Println("Unknown type")
		}
	}
	fmt.Println(stack)  // [21]
}