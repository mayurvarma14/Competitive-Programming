from typing import List


class Solution:
    def moveZeroes(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        write = 0
        for read, num in enumerate(nums):
            if num != 0:
                nums[write], nums[read] = nums[read], nums[write]
                write += 1


s = Solution()
nums = [0, 1, 0, 3, 12, 0, 42]
s.moveZeroes(nums)
print(nums)
