class Solution:
    def isValidSudoku(self, board: list[list[str]]) -> bool:
        rowSet = [set() for _ in range(9)]
        colSet = [set() for _ in range(9)]
        boxSet = [set() for _ in range(9)]
        
        for i in range(81):
            row, col = i // 9, i % 9
            c = board[row][col]
            if c == '.':  continue
            box = ((row//3)*3) + (col//3)

            if c in rowSet[row] or c in colSet[col] or c in boxSet[box]: 
                return False

            rowSet[row].add(c)
            colSet[col].add(c)
            boxSet[box].add(c)
        return True