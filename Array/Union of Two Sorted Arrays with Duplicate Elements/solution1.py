class Solution:

    # Function to return a list containing the union of the two arrays.
    def findUnion(self, a, b):
        a_length = len(a)
        b_length = len(b)
        union = [min(a[0], b[0])]
        read_a = read_b = 0
        write = 0
        while read_a < a_length and read_b < b_length:
            if a[read_a] == union[write]:
                read_a += 1
            elif b[read_b] == union[write]:
                read_b += 1
            elif a[read_a] <= b[read_b]:
                write += 1
                union.append(a[read_a])
                read_a += 1

            elif a[read_a] >= b[read_b]:
                write += 1
                union.append(b[read_b])
                read_b += 1

        for i in range(read_a, a_length):
            if union[write] != a[i]:
                union.append(a[i])
                write += 1
        for i in range(read_b, b_length):
            if union[write] != b[i]:
                union.append(b[i])
                write += 1
        return union


s = Solution()
print(s.findUnion([-7, 8], [-8, -3, 8]))
