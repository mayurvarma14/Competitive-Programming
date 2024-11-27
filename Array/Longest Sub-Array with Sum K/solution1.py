class Solution:
    def lenOfLongestSubarr(self, arr, k):
        pre_sum_map = {}
        total_sum = 0
        max_length = 0
        for i, num in enumerate(arr):
            total_sum += num
            if total_sum == k:
                max_length = max(max_length, i + 1)

            if pre_sum_map.get(total_sum - k) is not None:
                temp = i - pre_sum_map[total_sum - k]
                max_length = max(max_length, temp)

            if pre_sum_map.get(total_sum) is None:
                pre_sum_map[total_sum] = i
        return max_length


s = Solution()
print(s.lenOfLongestSubarr([2, 0, 0, 3], 3))
