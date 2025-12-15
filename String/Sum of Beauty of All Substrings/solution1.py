class Solution:
    def beautySum(self, s: str) -> int:
        n = len(s)
        count = 0
        for i in range(n):
            freq = {}
            for j in range(i, n):

                freq[s[j]] = freq.get(s[j], 0) + 1
                count += max(freq.values()) - min(freq.values())
        return count


s = Solution()
print(s.beautySum("aabcb"))  # Output: 5
