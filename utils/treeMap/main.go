package main

type treeNode struct {
	key   int
	val   int
	left  *treeNode
	right *treeNode
	// 记录，以该节点为根的 BST 有多少个节点
	size int
}

type TreeMap struct {
	root *treeNode
}

func NewTreeMap() *TreeMap {
	return &TreeMap{}
}
func (t *TreeMap) nodeSize(node *treeNode) int {
	if node == nil {
		return 0
	}
	return node.size
}

func (t *TreeMap) Size() int {
	return t.nodeSize(t.root)
}

func (t *TreeMap) IsEmpty() bool {
	return t.Size() <= 0
}

func (t *TreeMap) Get(key int) (int, bool) {
	x := t.get(t.root, key)
	if x == nil {
		// key 不存在
		return 0, false
	}

	return x.val, true
}

func (t *TreeMap) get(node *treeNode, key int) *treeNode {
	if node == nil {
		return nil
	}
	if node.key < key {
		return t.get(node.right, key)
	}
	if node.key > key {
		return t.get(node.left, key)
	}
	// node.key == key
	return node
}

// 添加 key -> val 键值对，如果键 key 已存在，则将值修改为 val
func (t *TreeMap) Put(key, val int) (int, bool) {
	oldVal, found := t.Get(key)
	t.root = t.put(t.root, key, val)
	return oldVal, found
}

func (t *TreeMap) put(node *treeNode, key, val int) *treeNode {
	if node == nil {
		return &treeNode{key: key, val: val, size: 1}
	}
	if node.key > key {
		node.left = t.put(node.left, key, val)
	}
	if node.key < key {
		node.right = t.put(node.right, key, val)
	}

	// node.key == key
	// 维护每个节点的 size 变量
	node.size = t.nodeSize(node.left) + t.nodeSize(node.right) + 1

	return node
}

// 判断 key 是否存在 Map 中
func (t *TreeMap) ContainsKey(key int) bool {
	return t.get(t.root, key) != nil
}

// 返回小于等于 key 的最大的键
func (t *TreeMap) FloorKey(key int) (int, bool) {
	if t.IsEmpty() {
		return 0, false
	}

	x := t.floorKey(t.root, key)
	if x == nil {
		return 0, false
	}
	return x.key, true
}

func (t *TreeMap) floorKey(node *treeNode, key int) *treeNode {
	if node == nil {
		return nil
	}
	// node.key > key，去左子树找
	if node.key > key {
		return t.floorKey(node.left, key)
	}
	// node.key < key，去右子树找
	if node.key < key {
		x := t.floorKey(node.right, key)
		if x == nil {
			return node
		}
		return x
	}

	// node.key == key
	return node
}

// 返回大于等于 key 的最小的键
func (t *TreeMap) CeilingKey(key int) (int, bool) {
	if t.IsEmpty() {
		return 0, false
	}

	x := t.ceilingKey(t.root, key)
	if x == nil {
		return 0, false
	}
	return x.key, true
}

func (t *TreeMap) ceilingKey(node *treeNode, key int) *treeNode {
	if node == nil {
		return nil
	}

	if node.key > key {
		x := t.ceilingKey(node.left, key)
		if x == nil {
			return node
		}
		return x
	}
	if node.key < key {
		return t.ceilingKey(node.right, key)
	}
	// node.key == key
	return node
}

func (t *TreeMap) Keys() []int {
	list := []int{}
	t.traverse(t.root, &list)
	return list
}

func (t *TreeMap) traverse(node *treeNode, list *[]int) {
	if node == nil {
		return
	}
	t.traverse(node.left, list)
	// 中序遍历
	*list = append(*list, node.key)
	t.traverse(node.right, list)
}
