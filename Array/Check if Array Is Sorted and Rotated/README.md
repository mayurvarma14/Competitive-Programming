# Check if Array Is Sorted and Rotated

## **Problem Summary**

- **Goal:** Determine if the given array `nums` can be obtained by rotating a non-decreasing sorted array.
- **Input:** An array `nums` of integers (`1 ≤ nums.length ≤ 100`, `1 ≤ nums[i] ≤ 100`).
- **Output:** `True` if `nums` is a rotation of a non-decreasing sorted array; otherwise, `False`.

## **Understanding the Problem**

- **Non-decreasing Order:** Each element is greater than or equal to the previous one.
- **Rotation:** Shifting elements to the right or left, wrapping around the ends.

### **Rotation Explanation**

- **Rotation by `x` positions:**
  - For array `A`, rotated array `B` satisfies `A[i] == B[(i + x) % A.length]`.
  - **Example:** Rotating `[1, 2, 3, 4, 5]` by 2 positions results in `[3, 4, 5, 1, 2]`.

- **Duplicates Allowed:** Elements in the array may repeat.

## **Examples**

### **Example 1**

- **Input:** `nums = [3, 4, 5, 1, 2]`
- **Output:** `True`
- **Explanation:**
  - Original sorted array: `[1, 2, 3, 4, 5]`
  - Rotated by 3 positions: `[3, 4, 5, 1, 2]`

### **Example 2**

- **Input:** `nums = [2, 1, 3, 4]`
- **Output:** `False`
- **Explanation:**
  - No rotation of a sorted array can result in `[2, 1, 3, 4]`.

### **Example 3**

- **Input:** `nums = [1, 2, 3]`
- **Output:** `True`
- **Explanation:**
  - Original sorted array: `[1, 2, 3]`
  - No rotation needed (rotation by 0 positions).

## **How to Determine if the Array is Sorted and Rotated**

1. **Identify Break Points:**
   - A break point occurs when a number is greater than its next number.
   - In a sorted array, there should be **no break points**.
   - In a rotated sorted array, there should be **at most one break point**.

2. **Check the Number of Break Points:**
   - **Zero Break Points:**
     - The array is sorted but not rotated (still valid).
   - **One Break Point:**
     - The array is rotated.
   - **More Than One Break Point:**
     - The array is not sorted and rotated.

3. **Edge Cases:**
   - If the last element is greater than the first element in a rotated array, ensure this does not create an extra break point.

## **Algorithm Steps**

1. **Initialize** a counter `count` to `0`.
2. **Loop** through the array `nums`:
   - For `i` from `0` to `len(nums) - 1`:
     - **Compare** `nums[i]` with `nums[(i + 1) % len(nums)]`:
       - If `nums[i] > nums[(i + 1) % len(nums)]`:
         - **Increment** `count` by `1`.
3. **Check** the `count`:
   - If `count` is `0` or `1`, **return** `True`.
   - Otherwise, **return** `False`.

## **Simple Code Example (Python)**

```python
def check(nums):
    count = 0
    n = len(nums)
    for i in range(n):
        # Compare current element with the next (circularly)
        if nums[i] > nums[(i + 1) % n]:
            count += 1
        # Early exit if more than one break point
        if count > 1:
            return False
    return True
```

## **Key Points to Remember**

- **At Most One Break Point:** A sorted and rotated array can have only one place where the order decreases.
- **Circular Comparison:** Use modulo `%` to compare the last element with the first.
- **Early Termination:** If you find more than one break point, you can return `False` immediately.
- **Handles No Rotation:** The algorithm works even if the array is not rotated (sorted in non-decreasing order).

## **Why Learn This?**

- **Array Manipulation Skills:** Enhances understanding of array traversal and circular arrays.
- **Problem-Solving Techniques:** Introduces the concept of break points in sequences.
- **Efficient Algorithms:** Teaches how to optimize by early exiting when conditions are met.
- **Coding Interviews:** Common type of question to test logical reasoning and coding proficiency.

---

**Remember:** Focus on counting the number of times the array order decreases when traversing. If it happens more than once, the array isn't a sorted array that has been rotated!