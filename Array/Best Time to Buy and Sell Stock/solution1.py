from typing import List


class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        max_profit = 0
        profit = 0
        buy = prices[0]
        for _, price in enumerate(prices):
            profit = price - buy
            max_profit = max(max_profit, profit)
            if profit < 0:
                profit = 0
                buy = price
        return max_profit


s = Solution()
print(s.maxProfit([7, 1, 5, 3, 6, 4]))
