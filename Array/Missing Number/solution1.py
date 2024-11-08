from typing import List


class Solution:
    def missingNumber(self, nums: List[int]) -> int:
        n = len(nums)
        total_sum = (n * (n + 1)) // 2
        current_sum = sum(nums)
        return total_sum - current_sum


s = Solution()
print(s.missingNumber([3, 0, 1]))
