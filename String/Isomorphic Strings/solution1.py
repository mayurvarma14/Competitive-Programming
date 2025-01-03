class Solution:
    def isIsomorphic(self, s: str, t: str) -> bool:
        mapping = {}
        mapped = set()
        for i, c in enumerate(s):

            if c not in mapping and t[i] not in mapped:
                mapping[c] = t[i]
                mapped.add(t[i])
            elif c not in mapping and t[i] in mapped:
                return False
            else:
                if mapping[c] != t[i]:
                    return False
        return True


s = Solution()
print(s.isIsomorphic("paper", "title"))
# print(s.isIsomorphic("badc", "baba"))
# print(s.isIsomorphic("bbbaaaba", "aaabbbba"))
