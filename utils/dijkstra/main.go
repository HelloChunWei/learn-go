package main

import (
	"container/heap"
	"math"
)

// Edge 代表一條邊：連到哪個節點、這條邊的權重
type Edge struct {
	To     *Node
	Weight int
}

type Node struct {
	// 編號
	Val       int
	Neighbors []Edge // 權重現在綁在邊上，不是節點上
}

type State struct {
	// 当前节点 ID
	node int
	// 从起点 s 到当前 node 节点的最小路径权重和
	distFromStart int
	Neighbors     []Edge
}

type PriorityQueue []*State

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].distFromStart < pq[j].distFromStart
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(*State))
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	res := old[n-1]
	*pq = old[:n-1]
	return res
}

// dijkstra 計算 src 到每個節點的最短距離
// size：圖上總共有幾個節點（假設節點編號落在 [0, size) 之間）
// src：起點
func dijkstra(size int, src *Node) []int {
	distance := make([]int, size)
	// 沒有計算前，全部都塞最大值
	for i := range distance {
		distance[i] = math.MaxInt32
	}
	pq := &PriorityQueue{}
	heap.Init(pq)
	// 起點到起點是 0
	distance[src.Val] = 0

	heap.Push(pq, &State{node: src.Val, distFromStart: 0, Neighbors: src.Neighbors})

	for pq.Len() > 0 {
		curState := heap.Pop(pq).(*State)
		node := curState.node
		dis := curState.distFromStart
		if distance[node] < dis {
			continue
		}
		for _, e := range curState.Neighbors {
			nextNode := e.To
			nextDis := dis + e.Weight
			if distance[nextNode.Val] <= nextDis {
				continue
			}
			heap.Push(pq, &State{node: nextNode.Val, distFromStart: nextDis, Neighbors: nextNode.Neighbors})
			distance[nextNode.Val] = nextDis
		}
	}
	return distance
}
