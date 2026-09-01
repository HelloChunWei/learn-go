package main

const R = 256

type TrieNode struct {
	Val      interface{}
	Children [R]*TrieNode
}

type TrieMap struct {
	Size int
	Root *TrieNode
}

func NewTrieMap() *TrieMap {
	return &TrieMap{Size: 0, Root: &TrieNode{}}
}

func (this *TrieMap) getNode(node *TrieNode, key string) *TrieNode {
	p := node
	for i := 0; i < len(key); i++ {
		// 沒找到
		if p == nil {
			return nil
		}
		c := key[i]
		p = p.Children[c]
	}
	return p
}

func (this *TrieMap) Put(key string, val interface{}) {
	if !this.ContainsKey(key) {
		this.Size++
	}
	this.Root = this._put(this.Root, key, val, 0)
}
func (this *TrieMap) _put(node *TrieNode, key string, val interface{}, i int) *TrieNode {
	if node == nil {
		return &TrieNode{}
	}
	if i == len(key) {
		// 找到了
		node.Val = val
		return node
	}
	c := key[i]
	node.Children[c] = this._put(node.Children[c], key, val, i+1)
	return node
}

func (this *TrieMap) Remove(key string) {
	if !this.ContainsKey(key) {
		return
	}
	this.Root = this._remove(this.Root, key, 0)
	this.Size--
}

func (this *TrieMap) _remove(node *TrieNode, key string, i int) *TrieNode {
	if node == nil {
		return nil
	}
	if i == len(key) {
		node.Val = nil
	} else {
		c := key[i]
		// 递归去子树进行删除
		node.Children[c] = this._remove(node.Children[c], key, i+1)
	}
	// 后序位置，递归路径上的节点可能需要被清理
	if node.Val != nil {
		// 如果该 TireNode 存储着 val，不需要被清理
		return node
	}
	// 检查该 TrieNode 是否还有后缀
	for c := 0; c < R; c++ {
		if node.Children[c] != nil {
			// 只要存在一个子节点（后缀树枝），就不需要被清理
			return node
		}
	}
	// 既没有存储 val，也没有后缀树枝，则该节点需要被清理
	return nil
}

func (this *TrieMap) Get(key string) interface{} {
	node := this.getNode(this.Root, key)
	if node == nil || node.Val == nil {
		return nil
	}
	return node.Val
}

func (this *TrieMap) ContainsKey(key string) bool {
	return this.Get(key) != nil
}

func (this *TrieMap) ShortestPrefixOf(key string) string {
	node := this.Root
	for i := 0; i < len(key); i++ {
		if node == nil {
			return ""
		}
		if node.Val != nil {
			return key[:i]
		}
		c := key[i]
		node = node.Children[c]
	}

	if node != nil && node.Val != nil {
		return key
	}
	return ""
}

func (this *TrieMap) LongestPrefixOf(key string) string {
	node := this.Root
	maxLen := 0
	for i := 0; i < len(key); i++ {
		if node == nil {
			break
		}
		if node.Val != nil {
			maxLen = i
		}
		c := key[i]
		node = node.Children[c]
	}

	if node != nil && node.Val != nil {
		return key
	}
	return key[:maxLen]
}

func (this *TrieMap) KeysWithPrefix(prefix string) []string {
	res := []string{}
	node := this.getNode(this.Root, prefix)
	// 代表沒有
	if node == nil {
		return res
	}
	// 找到 preFix 後，遍歷底下的所有child
	path := []byte(prefix)
	this.traverse(node, &path, &res)
	return res
}

func (this *TrieMap) traverse(node *TrieNode, path *[]byte, res *[]string) {
	if node == nil {
		return
	}
	if node.Val != nil {
		// 找到一個result
		*res = append(*res, string(*path))
	}

	for c := 0; c < R; c++ {
		*path = append(*path, byte(c))
		this.traverse(node.Children[c], path, res)
		// pop
		*path = (*path)[:len(*path)-1]
	}
}

func (this *TrieMap) HasKeyWithPrefix(prefix string) bool {
	node := this.getNode(this.Root, prefix)
	return node != nil
}

func (this *TrieMap) KeysWithPattern(pattern string) []string {
	res := []string{}
	path := []byte{}
	this.traversePattern(this.Root, &path, pattern, 0, &res)
	return res
}
func (this *TrieMap) traversePattern(node *TrieNode, path *[]byte, pattern string, i int, res *[]string) {
	if node == nil {
		return
	}
	if i == len(pattern) {
		if node.Val != nil {
			*res = append(*res, string(*path))
		}
		return
	}
	c := pattern[i]
	if c == '.' {
		// 代表可以配任何
		for j := 0; j < R; j++ {
			*path = append(*path, byte(j))
			this.traversePattern(node.Children[j], path, pattern, i+1, res)
			// pop
			*path = (*path)[:len(*path)-1]
		}
	} else {
		// 只能匹配 c
		*path = append(*path, byte(c))
		this.traversePattern(node.Children[c], path, pattern, i+1, res)
		// pop
		*path = (*path)[:len(*path)-1]
	}
}

// 這效率不好，因為是會找到全部的
func (this *TrieMap) HasKeyWithPattern(pattern string) bool {
	return len(this.KeysWithPattern(pattern)) > 0
}

func (this *TrieMap) _hasKeyWithPattern(node *TrieNode, pattern string, i int) bool {
	if node == nil {
		return false
	}
	if i == len(pattern) {
		return node.Val != nil
	}
	c := pattern[i]
	if c == '.' {
		for j := 0; j < R; j++ {
			if this._hasKeyWithPattern(node.Children[j], pattern, i+1) {
				return true
			}
		}
	} else {
		return this._hasKeyWithPattern(node.Children[c], pattern, i+1)
	}
	return false
}

func (this *TrieMap) Len() int {
	return this.Size
}

func main() {

}
