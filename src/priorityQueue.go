package main

import (
	"fmt"
	"strings"
	"strconv"
	"math"
)

// 優先度付きキューを実装する
// 今回は、最も大きい値を取り出すという優先度付きキューを実装する
// これは最大ヒープ（Maxヒープ）の完全二分木で実装できる
// 最大ヒープ（Maxヒープ）の完全二分木は、根（root）が常に最大である
// つまり、insertが起こるたびに、値を完全二分木の適切な位置に挿入する
// そして、extractが起こるたびに、その時点での完全二分木の根を値として取り出す
// 根の値を取り出したら、根が空きになるので、なんでも良いのだが、ここではキュー内の最後の要素を根に割り当て、その後それを完全二分木の適切な場所へ挿入する

var commands = [...]string{
	"insert 8",
	"insert 2",
	"extract",
	"insert 10",
	"extract",
	"insert 11",
	"extract",
	"extract",
	"end",
	
	// "insert 25",

	// "extract",
}

var priorityQueue []int
// var priorityQueue = []int{86, 14, 37, 12, 5, 32, 11, 7, 1, 2}
// var priorityQueue = []int{86, 37, 32, 12, 14, 25, 11, 7, 1, 2, 3}

func priorityQueuePractice() {
	for _, command := range commands {
		if strings.HasPrefix(command, "insert") {
			parts := strings.Split(command, " ")
			value, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("Failed to convert string to int:", err)
				continue
			}
			priorityQueue = append(priorityQueue, value)
			moveTargetKey := len(priorityQueue) - 1  // 新たにinsertした要素のキーを特定する
			lastWithUpComparisonAndMove(moveTargetKey)
		} else if command == "extract" {
			if len(priorityQueue) > 0 {
				extractValue := priorityQueue[0]
				priorityQueue[0] = priorityQueue[len(priorityQueue) - 1]  // キュー内の最後の要素を先頭に移動させる
				priorityQueue = priorityQueue[:len(priorityQueue)-1]  // キューの最後の要素は先頭へ移動させたので、最後の要素を含まないpriorityQueueに変化させる
				upperWithDownComparisonAndMove(0)
				fmt.Println(extractValue)
			} else {
				fmt.Println("キューの中は空なので、取り出せる値がありません")
			}
		} else {
			break
		}
	}
	fmt.Println(priorityQueue)
}

// キューの末から先頭へ（完全二分木における子から親方向へ）順に比較し、移動させる
func lastWithUpComparisonAndMove(moveTargetKey int) {
	if moveTargetKey == 0 {  // insertで指定された、新たにキューに追加した要素が根まで辿り着いた場合、これ以上比較対象がないので、ここで終了させる（再帰の基底条件）
		return
	}
	var parentKey int
	parentKey = (moveTargetKey - 1) / 2  // 新たに追加した要素が左の子であると仮定して、親要素のキーを求める
	if !isInteger(float64(parentKey)) {  // 親要素のキーが少数だった場合、新たに追加した要素は左の子ではなく右の子であったことが分かる
		// 右の子として親要素のキーを求める
		parentKey = (moveTargetKey - 2) / 2
	}
	if priorityQueue[moveTargetKey] > priorityQueue[parentKey] {
		priorityQueue[moveTargetKey], priorityQueue[parentKey] = priorityQueue[parentKey], priorityQueue[moveTargetKey]
		lastWithUpComparisonAndMove(parentKey)
	}
	return
}

// キューの先頭から末へ（完全二分木における親から子方向へ）順に比較し、移動させる
func upperWithDownComparisonAndMove(moveTargetKey int) {
	leftChildKey := (moveTargetKey * 2) + 1
	rightChildKey := (moveTargetKey * 2) + 2

	var compareTargetKey int

	if leftChildKey < len(priorityQueue) && rightChildKey < len(priorityQueue) {
		// ２つの子がいる場合
		if priorityQueue[leftChildKey] >= priorityQueue[rightChildKey] {  // ヒープには親子間での大小関係は存在するが、兄弟間での大小関係は存在しないので、左の子と右の子が同じ値なのであれば、ここでは左の子を比較対象とする
			compareTargetKey = leftChildKey
		} else {
			compareTargetKey = rightChildKey
		}
	} else if leftChildKey < len(priorityQueue) {  // 完全二分木なので、右の子のみいる状態はあり得ない。つまり左の子のみ存在しているケースだけを考えれば良い
		// 左の子だけいる場合
		compareTargetKey = leftChildKey
	} else {
		// 子がいない場合
		return
	}

	if priorityQueue[moveTargetKey] < priorityQueue[compareTargetKey] {  // 最大ヒープ（親が子より大きい状態）にしたいので、既にそのようになっているものに関しては何もする必要がない。つまり。親が子より小さい場合だけ、それを逆にしてあげる
		priorityQueue[moveTargetKey], priorityQueue[compareTargetKey] = priorityQueue[compareTargetKey], priorityQueue[moveTargetKey]
		upperWithDownComparisonAndMove(compareTargetKey)  // ここで再帰的に呼ぶ  compareTargetKeyを渡す  compareTargetKeyが次のmoveTargetKeyとして機能する
	}
	return
}

func isInteger(value float64) bool {
	// 引数で受け取った値と、その値の少数部分を切り捨てた値（整数部分のみ）を比較する
	// 同じであったら、引数で受け取った値は整数であったことが分かる
	// 異なった場合は、引数で受け取った値は少数であったことが分かる
	return value == math.Trunc(value)
}

// 8
// 10
// 11
// 2

// [86 25 37 12 14 32 11 7 1 2 5]

// 86
// [37 14 32 12 3 25 11 7 1 2]