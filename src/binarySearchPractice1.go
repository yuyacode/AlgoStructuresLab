package main

import (
    "fmt"
    "sort"
)

// weights : 降順に並んだ各荷物の重さ  [8 7 5 3 1]
// index : 荷物の重量が格納されたweightsの最大インデックス
// trucks : 擬似トラックのスライス  容量は3
// midWeight : 二分で探索する際に使用するmid
func canPartition(weights []int, index int, trucks []int, midWeight int) bool {
    if index < 0 {
        return true
    }
    // fmt.Println(trucks)
    for i := range trucks {  // i = 0, 1, 2
        if trucks[i] + weights[index] <= midWeight {  // トラックに荷物を軽い順で積んでいく  二分探索の分割点になっている最大重量の中央値以下である場合、まだまだそのトラックに荷物を積めるので積んでいく
            trucks[i] += weights[index]
            if canPartition(weights, index-1, trucks, midWeight) {  // 再帰的に呼ぶ
                // index-1により、次に重い荷物をそのトラックに積んでいく  この調子で１つのトラックに荷物を１つ１つ積んでいくと、どこかのタイミングで、トラックに積まれている荷物の重量が最大重量の中央値を超える
                // 超えた場合、forループにより次のトラックに残っている荷物を続きから積んでいく
                // そうしていくと、どこかで荷物が尽きる  そのときにindex < 0に条件が成立しtrueが返る
                // trueが返る限り、荷物を３つのトラックに正常に分配できたことになるので、より最大重量を下げることができないか調べる
                // falseが返る場合は、荷物を３つのトラックに分配できなかったことになるので、最大重量を上げることで対応できないか調べる
                return true
            }
            trucks[i] -= weights[index]
        }
        if trucks[i] == 0 {
            break
        }
    }
    return false
}

func findMinLoad(weights []int, numTrucks int) int {
    sum := 0  // 各荷物の合計の重さを格納する変数
    maxWeight := 0  // 重さが最大の荷物の重さを格納する変数
    for _, weight := range weights {
        sum += weight
        if weight > maxWeight {
            maxWeight = weight  // 重さが最大の荷物が見つかり次第、最大重量を更新していく
        }
    }
    min := maxWeight  // 荷物の最大重量を最小として扱う。これを下回るトラックの最大重量は、この重量が最大の荷物を積むことができないため、つまりトラックの最大重量の下限はmaxWeightになると判断できる
    max := sum  // トラックの最大重量の上限は、全荷物の合計重量になることが考えられる。全ての荷物を１つのトラックに積むことを想定する。これが最大積載量の上限となる
    for min < max {
        mid := (min + max) / 2
        trucks := make([]int, numTrucks)  // 擬似トラックを用意
        if canPartition(weights, len(weights)-1, trucks, mid) {
            max = mid  // 二分探索により前半を調べる
        } else {
            min = mid + 1  // 二分探索により後半を調べる
        }
    }
    return min  // 求められた最大積載量の最小値をreturn
}

func binarySearchPractice1() {
    weights := []int{8, 1, 7, 5, 3}  // 各荷物の重さ
    numTrucks := 3  // トラックの数
    sort.Sort(sort.Reverse(sort.IntSlice(weights)))  //  weightsを降順にソート [8 7 5 3 1]
    minLoad := findMinLoad(weights, numTrucks)
    fmt.Println("最小の最大積載量は:", minLoad)
}
