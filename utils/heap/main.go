package main

import (
	"fmt"
)

type SimpleMinPQ struct {
	arr  []int
	size int
}

func (pq *SimpleMinPQ) parent(node int) int {
	return (node - 1) / 2
}

func (pq *SimpleMinPQ) left(node int) int {
	return (node * 2) + 1
}

func (pq *SimpleMinPQ) right(node int) int {
	return (node * 2) + 2
}

func (pq *SimpleMinPQ) swap(i, j int) {
	pq.arr[i], pq.arr[j] = pq.arr[j], pq.arr[i]
}

func (pq *SimpleMinPQ) peak() int {
	return pq.arr[0]
}

func (pq *SimpleMinPQ) push(val int) {
	pq.arr = append(pq.arr, val)
	pq.swim(pq.size)
	pq.size++
}

func (pq *SimpleMinPQ) swim(node int) {
	parent := pq.parent(node)
	for node > 0 && pq.arr[parent] > pq.arr[node] {
		pq.swap(parent, node)
		node = parent
		parent = pq.parent(node)
	}
}

func (pq *SimpleMinPQ) pop() int {
	res := pq.arr[0]
	pq.arr[0] = pq.arr[pq.size-1]
	pq.size--
	pq.sink(0)
	return res
}

func (pq *SimpleMinPQ) sink(node int) {
	for {
		minNode := node
		left := pq.left(node)
		right := pq.right(node)

		// 因為要找最小，所以都要用 minNode
		if left < pq.size && pq.arr[minNode] > pq.arr[left] {
			minNode = left
		}

		if right < pq.size && pq.arr[minNode] > pq.arr[right] {
			minNode = right
		}

		// 代表左邊或右邊都比當前 node 大
		if minNode == node {
			break
		}
		pq.swap(minNode, node)
		node = minNode
	}
}

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x

}

func main() {
	pq := &SimpleMinPQ{
		[]int{},
		0,
	}

	pq.push(5)
	pq.push(3)
	pq.push(9)
	pq.push(10)
	pq.push(2)

	first := pq.pop()
	//2
	fmt.Printf("first: %d\n", first)

	sec := pq.pop()
	//3
	fmt.Printf("sec: %d\n", sec)
	thir := pq.pop()
	//5
	fmt.Printf("thir: %d\n", thir)

}
