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
