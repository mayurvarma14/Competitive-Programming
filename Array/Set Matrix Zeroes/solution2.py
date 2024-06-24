from typing import List


class Solution:
    def setZeroes(self, matrix: List[List[int]]) -> None:
        """
        Do not return anything, modify matrix in-place instead.
        """
        first_col = False
        rows = len(matrix)
        cols = len(matrix[0])

        for row_idx in range(rows):
            if matrix[row_idx][0] == 0:
                first_col = True
            for col_idx in range(1, cols):
                if matrix[row_idx][col_idx] == 0:
                    matrix[row_idx][0] = 0
                    matrix[0][col_idx] = 0

        for row_idx in range(rows - 1, -1, -1):
            for col_idx in range(cols - 1, 0, -1):
                if matrix[row_idx][0] == 0 or matrix[0][col_idx] == 0:
                    matrix[row_idx][col_idx] = 0
            if first_col is True:
                matrix[row_idx][0] = 0
        return matrix


s = Solution()
print(s.setZeroes([[0, 1, 2, 0], [3, 4, 5, 2], [1, 3, 1, 5]]))
