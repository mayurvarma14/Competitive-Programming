from typing import List


class Solution:
    def findMaxConsecutiveOnes(self, nums: List[int]) -> int:
        max_con = count = 0
        for num in nums:
            if num == 0:
                max_con = max(max_con, count)
                count = 0
            else:
                count += 1
        max_con = max(max_con, count)
        return max_con


s = Solution()
print(s.findMaxConsecutiveOnes([1, 1, 0, 1, 1, 1]))
