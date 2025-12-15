class Solution:
    def romanToInt(self, s: str) -> int:
        mapping = {"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
        result = 0
        prev = ""
        for r in s:
            if (
                (prev == "I" and r in ["V", "X"])
                or (prev == "X" and r in ["L", "C"])
                or (prev == "C" and r in ["D", "M"])
            ):
                result += mapping[r] - mapping[prev] * 2
            else:
                result += mapping[r]
            prev = r
        return result


s = Solution()
print(s.romanToInt("MCMXCIV"))
