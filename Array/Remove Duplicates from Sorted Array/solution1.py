from typing import List


class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        write = 0
        for read in range(1, len(nums)):
            if nums[read] != nums[write]:
                write += 1
                nums[write] = nums[read]
        return write + 1


s = Solution()
print(s.removeDuplicates([1, 1, 2, 3, 3, 4, 4, 5]))
