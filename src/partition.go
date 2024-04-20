package main

import (
	"fmt"
)

func partition() {
	array := []int{13, 19, 9, 5, 12, 8, 7, 4, 21, 2, 6, 11}
	comparingElements := array[len(array)-1]  //  11
	smallRangeTrailingKey := -1
	for key, value := range array {
		if value <= comparingElements {
			smallRangeTrailingKey++
			array[smallRangeTrailingKey], array[key] = value, array[smallRangeTrailingKey]
		}
	}
	fmt.Println(array)  // [9 5 8 7 4 2 6 11 21 13 19 12]
}