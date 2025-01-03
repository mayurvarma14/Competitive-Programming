class Solution:
    def largestOddNumber(self, num: str) -> str:
        res = ""
        for i, v in enumerate(reversed(num)):
            if int(v) % 2 != 0:
                return num[: len(num) - i]
        return res


s = Solution()
print(s.largestOddNumber("35427"))
