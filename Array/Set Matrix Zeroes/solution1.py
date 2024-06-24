from typing import List


class Solution:
    def setZeroes(self, matrix: List[List[int]]) -> None:
        """
        Do not return anything, modify matrix in-place instead.
        """
        row_zeroes = set()
        col_zeroes = set()
        for row_idx, row in enumerate(matrix):
            for col_idx, col in enumerate(row):
                if col == 0:
                    row_zeroes.add(row_idx)
                    col_zeroes.add(col_idx)
        for row_idx, row in enumerate(matrix):
            for col_idx, col in enumerate(row):
                if col_idx in col_zeroes or row_idx in row_zeroes:
                    matrix[row_idx][col_idx] = 0
        return matrix


s = Solution()
print(s.setZeroes([[1, 1, 1], [1, 0, 1], [1, 1, 1]]))
