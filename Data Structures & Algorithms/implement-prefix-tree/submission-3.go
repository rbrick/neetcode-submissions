
type node struct {
	children [26]*node
	end bool
}


type PrefixTree struct {
	root *node
}

func Constructor() PrefixTree {
    return PrefixTree{
		root: &node{
			children: [26]*node{},
		},
	}
}

func (this *PrefixTree) Insert(word string) {
	curr := this.root

	for _, r := range word {
		d := r - 'a'
		n := curr.children[d]

		if n == nil {
			n = &node{
				children: [26]*node{},
			}

			curr.children[d] = n
		}

		curr = n
	}

	curr.end = true
}

func (this *PrefixTree) Search(word string) bool {
	curr := this.root

	for _, r := range word {
		d := r - 'a'
		n := curr.children[d]

		// early return = full word was not in trie
		if n == nil {
			return false
		}

		curr = n
	}

	return curr.end
}

func (this *PrefixTree) StartsWith(prefix string) bool {
	curr := this.root

	for _, r := range prefix {
		d := r - 'a'
		n := curr.children[d]

		// early return = full word was not in trie
		if n == nil {
			return false
		}

		curr = n
	}

	return true
}
