package main

import (
	"fmt"
)

func recursionPractice1() {
	data := []int{1, 2, 3, 4, 5}
	minKey := 0
	maxKey := len(data)
	maxVal := findMaxValue(data, minKey, maxKey)
	fmt.Println(maxVal)
}

func findMaxValue(data []int, minKey int, maxKey int) int {
	midKey := (minKey + maxKey) / 2
	if minKey == maxKey - 1 {
		return data[minKey]
	} else {
		u := findMaxValue(data, minKey, midKey)
		fmt.Println(u)
		v := findMaxValue(data, midKey, maxKey)
		fmt.Println(v)
		x := maxInt(u, v)
		return x
	}
}

func maxInt(x, y int) int {
    if x > y {
        return x
    }
    return y
}
