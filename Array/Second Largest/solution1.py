class Solution:
    def getSecondLargest(self, arr):
        largest = arr[0]
        second_largest = -1
        for num in arr:

            if num > largest:
                second_largest = largest
                largest = num

            if num > second_largest and num != largest:
                second_largest = num
        return second_largest


s = Solution()
print(s.getSecondLargest([28078, 19451, 935, 28892, 2242, 3570, 5480, 231]))
