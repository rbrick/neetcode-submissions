class Solution:
    def exist(self, board: List[List[str]], word: str) -> bool:
        # position in word
    
        directions = [
            (1,0), # DOWN = y+1, x
            (-1, 0), # UP = y-1, x
            (0, 1), # RIGHT = y, x+1
            (0, -1) # LEFT
        ]
        
        def visit(built_word, cell, visited) -> bool:
            if built_word == word: return True    
            y, x = cell[0], cell[1]
            
            # out of bounds
            if y < 0 or y >= len(board): return False
            if x < 0 or x >= len(board[y]): return False
                
            at = board[y][x]

            if at != word[len(built_word)]:return False
            if built_word+at == word: return True
        
            visited.add(cell)
            # we can continue
            for d in directions:
                new_cell = (y+d[0], x+d[1])
                if new_cell in visited: continue 
                visited.add(new_cell)
                if visit(built_word+at, new_cell, visited): return True
                visited.remove(new_cell)
           
            # visited.remove(cell)
            return False
            
            
        visited = set()
        # initial loop
        for y in range(len(board)):
            for x in range(len(board[y])):
                cell = board[y][x]
                
                if cell == word[0]:
                    # visit surrounding cells
                    visited.add((y, x))
                    for d in directions:
                        new_cell= (y+d[0], x+d[1])
                        visited.add(new_cell)
                        if visit(built_word=cell, cell=new_cell, visited=visited): return True
                        visited.remove(new_cell)
                    visited.remove((y, x))
        return False