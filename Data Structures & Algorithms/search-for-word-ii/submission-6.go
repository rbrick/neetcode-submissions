
type TrieNode struct {
	children [26]*TrieNode
	end      bool
}

func (t *TrieNode) hasChild(c byte) bool {
	idx := c - 'a'
	return t.children[idx] != nil
}

func (t *TrieNode) getChild(c byte) *TrieNode {
	idx := c - 'a'
	return t.children[idx]
}

func newTrieNode() *TrieNode {
	return &TrieNode{
		children: [26]*TrieNode{},
	}
}

type Trie struct {
	root *TrieNode
}

func (t *Trie) insert(s string) {
	curr := t.root

	for _, c := range s {
		idx := c - 'a' // 0-25

		if curr.children[idx] == nil {
			curr.children[idx] = newTrieNode()
		}
		// update current node to child
		curr = curr.children[idx]
	}

	curr.end = true
}

func (t *Trie) hasWord(word string) bool {

	curr := t.root

	for _, c := range word {
		idx := c - 'a' // 0-25

		if curr.children[idx] == nil {
			return false
		}

		// update current node to child
		curr = curr.children[idx]
	}

	return curr.end
}

func newTrie() *Trie {
	return &Trie{
		root: newTrieNode(),
	}
}

func visitCell(trieNode *TrieNode, board [][]byte, row, col int, currWord string, words map[string]bool, visited [][]bool) map[string]bool {
	// out of bounds
	if row < 0 || row >= len(board) {
		return words
	}

	// out of bounds
	if col < 0 || col >= len(board[row]) {
		return words
	}

	if trieNode == nil {
		return words
	}

	// if visited[row][col] {
	// 	return words
	// }

	directions := [][]int{
		{1, 0},  // move down (increase row)
		{0, 1},  // move left (increase col)
		{-1, 0}, // move up (decrease row)
		{0, -1}, // move right (decrease col)
	}

	char := board[row][col]

	word := currWord + string(char)

	// fmt.Println(word)

	if trieNode.end {
		words[word] = true
	}

	for _, dirs := range directions {
		r, c := row+dirs[0], col+dirs[1]

		if r < 0 || r >= len(board) || c < 0 || c >= len(board[r]) {
			continue
		}

		if visited[r][c] {
			continue
		}

		visited[r][c] = true

		child := trieNode.getChild(board[r][c])

		for k, v := range visitCell(child, board, r, c, word, words, visited) {
			words[k] = v
		}

		visited[r][c] = false

	}

	return words
}

func findWords(board [][]byte, words []string) []string {
	trie := newTrie()

	for _, word := range words {
		trie.insert(word)
	}

	// visited := make([][]bool, len(board))

	// // initialize visited
	// for i, _ := range visited {
	// 	visited[i] = make([]bool, len(board[i]))
	// }

	/*
	  ["a","b","c","d"],
	  ["s","a","a","t"],
	  ["a","c","k","e"],
	  ["a","c","d","n"]
	*/

	foundWords := map[string]bool{}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {

			set := make(map[string]bool)

			visited := make([][]bool, len(board))

			for i := 0; i < len(visited); i++ {
				visited[i] = make([]bool, len(board[i]))
			}

			visited[i][j] = true

			words := visitCell(trie.root.getChild(board[i][j]), board, i, j, "", set, visited)

			for k := range words {
				foundWords[k] = true
			}
		}
	}

	arr := []string{}

	for k := range foundWords {
		arr = append(arr, k)
	}

	return arr
}
