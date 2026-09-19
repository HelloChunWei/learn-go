package main

type UF struct {
	// 有多少連通量
	count int
	// x 的 parent 就是 parent[x]
	parent []int
}

func NewUF(n int) *UF {
	parent := make([]int, n)
	for index := range parent {
		parent[index] = index
	}
	return &UF{
		count:  n,
		parent: parent,
	}
}
func (u *UF) Find(x int) int {
	if x != u.parent[x] {
		u.parent[x] = u.Find(u.parent[x])
	}
	return u.parent[x]
}

func (u *UF) Union(p, q int) {
	parentP := u.Find(p)
	parentQ := u.Find(q)
	if parentP == parentQ {
		return
	}
	u.parent[parentP] = parentQ
	u.count--
}

func (u *UF) Connect(p, q int) bool {
	parentP := u.Find(p)
	parentQ := u.Find(q)
	return parentP == parentQ
}

func (u *UF) Count() int {
	return u.count
}
