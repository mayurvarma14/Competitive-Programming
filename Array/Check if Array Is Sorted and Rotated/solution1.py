from typing import List


class Solution:
    def check(self, nums: List[int]) -> bool:
        pivot = 0
        length = len(nums)
        for i in range(length):
            if nums[i] > nums[(i + 1) % length]:
                pivot += 1
            if pivot > 1:
                return False
        return True


s = Solution()
print(s.check([2, 1, 3, 4]))
