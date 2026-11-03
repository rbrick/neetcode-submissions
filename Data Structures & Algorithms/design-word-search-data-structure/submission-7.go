
type node struct {
	children [26]*node
	end      bool
}

func newNode() *node {
	return &node{
		children: [26]*node{},
	}
}

type WordDictionary struct {
	root *node
}

func Constructor() WordDictionary {
	return WordDictionary{
		root: newNode(),
	}
}

func (this *WordDictionary) AddWord(word string) {
	curr := this.root

	for _, r := range word {
		idx := r - 'a'
		n := curr.children[idx]

		if n == nil {
			n = newNode()
			curr.children[idx] = n
		}

		curr = n
	}

	curr.end = true
}

func (this *WordDictionary) dfs(word string, depth int, n *node) bool {
	if n == nil {
		return false
	}

	if len(word) == depth {
		return n.end
	}

	c := word[depth]

	if c == '.' {
		for _, child := range n.children {
			if child != nil && this.dfs(word, depth+1, child) {
				return true
			}
		}

		return false
	}

	idx := c - 'a'
	return this.dfs(word, depth+1, n.children[idx])
}

func (this *WordDictionary) Search(word string) bool {
	return this.dfs(word, 0, this.root)
}
