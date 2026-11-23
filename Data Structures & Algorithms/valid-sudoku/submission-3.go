const (
	Width      = 9
	Height     = 9
	TotalCells = Width * Height
)

func isValidSudoku(board [][]byte) bool {
	var rowSet, colSet, boxSet [9][9]bool

	for i := 0; i < TotalCells; i++ {
		row := (i / Height)
		col := i % Width
		c := board[row][col]
		if c == '.' {
			continue
		}

		box := ((row/3)*3 + (col / 3))

		idx := c - '1'

		if rowSet[row][idx] || colSet[col][idx] || boxSet[box][idx] {
			return false
		}

		rowSet[row][idx] = true
		colSet[col][idx] = true
		boxSet[box][idx] = true

	}

	return true
}