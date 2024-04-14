package main

import (
	"fmt"
)

// processInfoをQueue（待ち行列）のように見立てている
func queue() {
	q := 100  // クオンタム(1回の処理能力に相当)
	cumulativeProcessingTime := 0  // 累計処理時間
	processInfo := map[string]int{"p1": 150, "p2": 80, "p3": 200, "p4": 350, "p5": 20}  // 処理するプロセス情報  プロセス名: 必要処理時間
	completedProcess := make(map[string]int)  // 処理が完了したプロセスを格納する場所
	for len(processInfo) > 0 {
		for processName, processTime := range processInfo {  // processTime：そのプロセスの処理にかかる時間
			if processTime > q {
				processTime -= q
				processInfo[processName] = processTime  // 再度処理する場合は、末尾に追加
				cumulativeProcessingTime += q  // 処理時間を加算
			} else {
				cumulativeProcessingTime += processTime  // 処理時間を加算
				completedProcess[processName] = cumulativeProcessingTime
				delete(processInfo, processName)  // 処理が完了した要素は削除
			}
		}
	}
	fmt.Println(completedProcess)
	// 出力は下記
	// Goのmapは順序を保証しないので、処理が上手くいっているのか認識しずらいが、上手くいっている
	// map[p1:450 p2:180 p3:550 p4:800 p5:400]
	
	// completedProcessにappendされた順序に変換すると下記
	// map[
	// 	p2:180 
	// 	p5:400
	// 	p1:450 
	// 	p3:550 
	// 	p4:800 
	// ]
}