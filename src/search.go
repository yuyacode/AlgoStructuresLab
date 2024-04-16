package main

import (
	"fmt"
)

// --------------------
// 線形探索
// --------------------

// 下記コードにおける最悪の場合の時間複雑性は O(n*m)
// １つの配列内から１つのデータを線形探索で探索する場合、最悪の場合の時間複雑性は O(n) nは配列の要素数となる

func linearSearch() {
	// 下記２つの数列に共通して存在する要素を線形探索で調べる
	numberSequenceS := []int{1, 2, 3, 4, 5}
	numberSequenceQ := []int{3, 4, 1}

	count := 0
	for _, seqQVal := range numberSequenceQ {
		for _, seqSVal := range numberSequenceS {
			if seqSVal == seqQVal {
				count++
				break  // 数列内の要素に重複はないという前提なので、1回条件が成立したら、これ以上は条件が成立しない（これ以上同じ要素はない）ことになるので、breakで抜けてしまう
			}
		}
	}
	fmt.Println(count)  // 3
}


// --------------------
// 二分探索
// --------------------
func binarySearch() {
	numberSequenceS := []int{4, 5, 6, 7, 7, 9, 10, 11, 12}  // 9個の要素
	numberSequenceQ := []int{3, 5, 40, 6, 8}
	count := 0
	for _, q := range numberSequenceQ{
		low, high := 0, len(numberSequenceS) - 1
		for low <= high {  // highとlowが逆転した瞬間に、数列Qに存在する該当の要素は数列Sには存在しないと判断できる。なので、それが成立するまでは二分で探索してあげる
			middle := (low + high) / 2
			if q == numberSequenceS[middle] {
				fmt.Println("一致")
				count++
				break  // 成立したら抜けないと無限ループが発生する
			} else if q < numberSequenceS[middle] {
				fmt.Println("まだまだ小さい")
				high = middle - 1
			} else {
				fmt.Println("まだまだ大きい")
				low = middle + 1
			}
		}
	}
	fmt.Println(count)  // 2
}



// --------------------
// ハッシュ
// --------------------
type hashTable struct {
	keys []int
	values []string
	size int
}

func initHashTable(size int) *hashTable {
	return &hashTable{
		keys: make([]int, size),
		values: make([]string, size),
		size: size
	}
}

func (h *hashTable) hashFunction(key int) {
	return key % h.size  // 今回はkeyをハッシュテーブルのサイズである10で割ったときの余りをハッシュ値とする
}

func (h *hashTable) insert(key int, value string) {
	index := h.hashFunction(key)
	for h.values[index] != "" {
		index = (index + 1) % h.size
		if index == h.hashFunction(key) {
			fmt.Println("Hash Table is full")  // 1周してきたら、もうハッシュテーブルは空いていないことになるので、満タンであると返す
			return
		}
	}
	h.keys[index] = key
	h.values[index] = value
}

func (h *hashTable) search(key int) (string, bool) {
	index := h.hashFunction(key)
	for h.values[index] != "" {
		if h.keys[index] == key {
			return h.values[index], true
		}
		index = (index + 1) % h.size  // insert時のプロービング（衝突が起きた場合に代替の空き位置を効率的に見つけるプロセス）と同じ方法で探索する
		if index == h.hashFunction(key) {  // 1周したらbreakで抜ける  全て探索し尽くして、もう確定で見つからないため
			break
		}
	}
	return "", false
}

func (h *hashTable) display() {
	for i, v := range h.values {
		if v != "" {
			fmt.Printf("key: %d, value: %s\n", h.keys[i] ,v)
		}
	}
}

// 値から生成したハッシュ値をキーとし、配列等の格納場所に格納することで、検索や削除の際に対象の要素から同じようにハッシュ値を生成し、得られたハッシュ値をキーに持つ要素にアクセスすることで、高速な検索や削除を実現する探索アルゴリズムの１つ
// O(n)とされている
func hash() {
	hashTable := initHashTable(10)
	hashTable.insert(1, "Apple")  // index 1
	hashTable.insert(2, "Banana")  // index 2
	hashTable.insert(11, "Cherry")  // index 3

	value, exists := hashTable.search(11)
	if exists {
		fmt.Println("found value : ", value)
	} else {
		fmt.Println("value not found")
	}

	hashTable.display()
}



// 生成されたハッシュ値の重複により発生する衝突を回避する方法としては、下記３つが挙げられる
// ・ チェイニング （Chaining）
// ・　オープンアドレッシング （Open Addressing）
// ・ 二次ハッシュ (Double Hashing)



// --------------------
// チェイニング （Chaining） のコード例
// --------------------
type ListNode struct {
    key   int
    value string
    next  *ListNode
}

type HashTable struct {  // ハッシュテーブルを示す構造体
    buckets []*ListNode
}
// type HashTable []*ListNode
// 上記のように、bucketsフィールドを経由せずに扱うことも可能
// この定義を行った場合、コードの煩雑性が低下し、可読性や読解性が向上するというメリットが存在する
// しかし、将来的にフィールドを持たせたい場合は、新たな型を定義する必要があるといったデメリットも存在する
// そのときの要件に応じて使い分けるべき

// 指定したサイズのハッシュテーブルを生成
func NewHashTable(size int) *HashTable {
    return &HashTable{buckets: make([]*ListNode, size)}
}

// ハッシュテーブルへの挿入を示したコード
func (h *HashTable) Put(key int, value string) {
    index := hashFunction(key, len(h.buckets))  // ハッシュ値を生成  これが後にハッシュテーブルのキー（つまり要素の格納場所）になる
    newNode := &ListNode{key: key, value: value}
    if h.buckets[index] == nil {
        h.buckets[index] = newNode
    } else {  // 同じハッシュ値を生成する要素が過去に存在し、衝突が発生した場合
        current := h.buckets[index]
        for current.next != nil {
            current = current.next
        }
        current.next = newNode  // 空いている場所が見つかったら、そこにノードを挿入
    }
}



// --------------------
// ・　オープンアドレッシング （Open Addressing） のコード例
// --------------------
func (h *HashTable) Put(key int, value string) {
    index := hashFunction(key, len(h.buckets))
	// これより下のコードは、あくまで空いている位置（index）を見つけるための手段の１つでしかなく、下記のコードに深い意味はない
	// ただ単に、空いている位置が見つかればそれで良いので、この方法以外に良い方法が個人で思いつくのであれば、またチームの決まりがあるなら、それに従うべき
    for i := 0; i < len(h.buckets); i++ {
        newIndex := (index + i) % len(h.buckets)  // 線形探索
        if h.buckets[newIndex] == nil {
            h.buckets[newIndex] = &ListNode{key: key, value: value}
            return
        }
    }
}



// --------------------
// ・ 二次ハッシュ (Double Hashing) のコード例
// --------------------
func doubleHashFunction1(key int) int {
    // 何らかのハッシュ計算
}

func doubleHashFunction2(key int) int {
    // 2つ目のハッシュ計算
}

func (h *HashTable) Put(key int, value string) {
    index := doubleHashFunction1(key)
    stepSize := doubleHashFunction2(key)
    for i := 0; i < len(h.buckets); i++ {
        newIndex := (index + i*stepSize) % len(h.buckets)
        if h.buckets[newIndex] == nil {
            h.buckets[newIndex] = &ListNode{key: key, value: value}
            return
        }
    }
}



// 補足：リンクリスト（Linked List）とは
// ・リンクリストは、データ要素がコンテナ内で直接的には連続していなくても良いという特性を持つ動的データ構造
// ・各要素（ノード）は自身のデータと次（または前）の要素の参照（リンク）を持っている

// 特徴と用途
// 動的なサイズ変更：リンクリストは実行時にサイズが変更可能。新しい要素の追加や既存要素の削除が容易で、それに伴うメモリ再配置のコストがかからないため、柔軟に扱うことができる
// データの挿入と削除の効率性：任意の位置におけるデータの挿入や削除が、ポインタの変更だけで完了するため、操作が高速（ただし、挿入/削除位置を見つけるためには前から順に走査する必要があるため、その点はO(n)の時間がかかる）。

// 種類
// ・単方向リンクリスト：各ノードが次のノードへのリンクのみを持つ
// ・双方向リンクリスト：各ノードが次のノードと前のノードへのリンクを両方持つ。これにより、前後どちらにも走査が可能
// ・循環リンクリスト：最後のノードが最初のノードを指し示す形でリンクされる