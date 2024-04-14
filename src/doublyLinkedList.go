package main

import "fmt"

// 連結リスト全体を表す型
type list struct {
	head  *node  // 先頭のノードを示すポインタ
	last  *node  // 最後のノードを示すポインタ
}

type node struct {
	value int
	prev *node  // 1つ前のノードを指し示すポインタ
	next *node  // 1つ後ろのノードを指し示すポインタ
}

func doublyLinkedList() {
	list := list{}
	list.print()
	list.append(1)
	list.print()
	list.append(2)
	list.print()
	list.append(3)
	list.print()
	list.append(4)
	list.print()
	list.remove()
	list.print()
	list.remove()
	list.print()
}

// appendは必ず先頭から行われるという前提
func (list *list) append(value int) {
	node := node{value: value}  // まずノードを作る
	if list.head == nil {  // リストの先頭が存在しない場合、リストは空と判断できる。この場合、リストの先頭に今回のノードを追加するだけ
		// 現在リストには、今回appendするノードしかないので、ノードのprevやnextといったフィールドをいじる必要はない
		list.head = &node  // リストの先頭に、今回appendするノードを指定
		list.last = &node  // リストの末尾に、今回appendするノードを指定
	} else {  // リストの先頭がnilではない場合、少なくとも１つはリスト内にノードが存在するので、そのノードを２番目としつつ、今回appendするノードを１番目にする
		oldHead := list.head  // 今先頭にあるノードをoldHeadに退避させる
		list.head = &node  // 先頭が空いたので、appendするノードを指定
		list.head.next = oldHead  // appendしたノードのnextに、元の先頭を指定することで、元の先頭を２番目に移す
		oldHead.prev = &node  // 元の先頭のprevに、appendしたノードを指定
	}
}

// removeは必ず先頭から行われるという前提
func (list *list) remove() {
	list.head = list.head.next
	list.head.prev = nil
}

func (list *list) print() {
	current := list.head
	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}
	fmt.Println("")

	// 出力
	// 1
	// 
	// 2
	// 1
	// 
	// 3
	// 2
	// 1
	// 
	// 4
	// 3
	// 2
	// 1
	// 
	// 3
	// 2
	// 1
	// 
	// 2
	// 1
}