class Solution:
    def removeOuterParentheses(self, s: str) -> str:
        stack = []
        res = ""
        for c in s:
            if c == "(":
                if len(stack) >= 1:
                    res += c
                stack.append(c)
            elif c == ")":
                stack.pop()
                if len(stack) >= 1:
                    res += c
        return res


s = Solution()
print(s.removeOuterParentheses("(()())(())(()(()))"))
