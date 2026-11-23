
func isValidSudoku(board [][]byte) bool {

	rowSet, colSet := make([][]bool, 9), make([][]bool, 9)

	for x := 0; x < 9; x++ {
		rowSet[x] = make([]bool, 9)
		colSet[x] = make([]bool, 9)
	}

	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			my := y * 3
			mx := x * 3

			boxSeen := make([]bool, 9)
			for nx := 0; nx < 3; nx++ {
				for ny := 0; ny < 3; ny++ {
					row := my + nx
					col := mx + ny
					c := board[row][col]
					if c == '.' {
						continue
					}

					idx := c - '1'

					if rowSet[row][idx] || colSet[col][idx] || boxSeen[idx] {
						return false
					}

					rowSet[row][idx] = true
					colSet[col][idx] = true
					boxSeen[idx] = true
				}
			}
		}
	}

	return true
}