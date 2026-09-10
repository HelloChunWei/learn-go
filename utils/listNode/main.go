package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseN(list *ListNode, n int) *ListNode {
	if list == nil || n <= 0 {
		return list
	}
	var pre *ListNode
	cur := list

	for n > 0 && cur != nil {
		next := cur.Next // 先存下一個
		cur.Next = pre   // 反轉指針
		pre = cur        // pre 往前
		cur = next       // cur 往前
		n--
	}

	list.Next = cur

	return pre
}

func reverse(list *ListNode) *ListNode {
	var pre *ListNode
	cur := list

	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func createList(vals []int) *ListNode {
	dummy := &ListNode{
		-1,
		nil,
	}

	head := dummy

	for _, val := range vals {
		node := &ListNode{
			val,
			nil,
		}
		head.Next = node
		head = head.Next
	}

	return dummy.Next

}

func printList(list *ListNode) {
	head := list
	for head != nil {
		fmt.Printf("val: %d\n", head.Val)
		head = head.Next
	}
}

// LRU cache
type NodeWithPrev struct {
	key, val   int
	prev, next *NodeWithPrev
}

func NewNodeWithPrev(key, val int) *NodeWithPrev {
	return &NodeWithPrev{key: key, val: val}
}

type DoubleList struct {
	size       int
	head, tail *NodeWithPrev
}

func NewDoubleList() *DoubleList {
	head := NewNodeWithPrev(0, 0)
	tail := NewNodeWithPrev(0, 0)
	head.next = tail
	tail.prev = head
	return &DoubleList{
		size: 0,
		head: head,
		tail: tail,
	}
}

func (this *DoubleList) addLast(x *NodeWithPrev) {
	x.prev = this.tail.prev
	x.next = this.tail
	this.tail.prev.next = x
	this.tail.prev = x
	this.size++
}

// x 一定在 list 中
func (this *DoubleList) remove(x *NodeWithPrev) {
	x.prev.next = x.next
	x.next.prev = x.prev
	this.size--
}

func (this *DoubleList) removeFirst() *NodeWithPrev {
	// 空的
	if this.head.next == this.tail {
		return nil
	}
	first := this.head.next
	this.remove(first)
	return first
}

func (this *DoubleList) Size() int {
	return this.size
}

type LruCache struct {
	_map  map[int]*NodeWithPrev
	cache *DoubleList
	cap   int
}

func NewLruCache(cap int) *LruCache {
	return &LruCache{
		_map:  make(map[int]*NodeWithPrev),
		cache: NewDoubleList(),
		cap:   cap,
	}
}

// 簡單封裝

// 把某個 key 變成最常用
func (this *LruCache) makeRecently(key int) {
	node := this._map[key]
	this.cache.remove(node)
	this.cache.addLast(node)
}

// 新增最常用
func (this *LruCache) addRecently(key, val int) {
	node := NewNodeWithPrev(key, val)
	this.cache.addLast(node)
	this._map[key] = node
}

func (this *LruCache) deleteKey(key int) {
	node := this._map[key]
	this.cache.remove(node)
	delete(this._map, key)
}

func (this *LruCache) removeLeastRecently() {
	// remove head
	first := this.cache.removeFirst()
	delete(this._map, first.key)
}

func (this *LruCache) Get(key int) int {
	if _, ok := this._map[key]; !ok {
		return -1
	}
	this.makeRecently(key)
	return this._map[key].val
}

func (this *LruCache) Put(key, val int) {
	// 如果有的話，update val ，然後讓他變成最新的
	if _, ok := this._map[key]; ok {
		node := this._map[key]
		node.val = val
		this.makeRecently(key)
		return
	}
	// 沒有的話，新增 node
	if this.cache.Size() >= this.cap {
		this.removeLeastRecently()
	}
	// 塞入
	this.addRecently(key, val)
}

func main() {
	list1 := createList([]int{1, 2, 3, 4, 5})
	fmt.Print("---- list 1 ---- \n")
	printList(list1)

	list2 := createList([]int{1, 2, 3, 4, 5}) // 獨立的 list
	list2 = reverse(list2)
	fmt.Print("--- after reverse --- \n")
	printList(list2)

	list3 := ReverseN(list1, 4) // list1 還是完整的
	fmt.Print("--- after reverseN n = 4 --- \n")
	printList(list3)

}
