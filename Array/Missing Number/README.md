# Missing Number

## **Problem Summary**

- **Goal:** Find the missing number in an array containing `n` distinct numbers taken from the range `0` to `n`.
- **Input:** An integer array `nums` of length `n`, containing unique numbers in the range `[0, n]`.
- **Output:** The single number in the range `[0, n]` that is not present in `nums`.

**Constraints:**

- `n == nums.length`
- `1 ≤ n ≤ 10⁴`
- `0 ≤ nums[i] ≤ n`
- All elements in `nums` are **unique**.

## **Understanding the Problem**

- **Array of Size `n`:** Contains numbers from `0` to `n`, but one number is missing.
- **Objective:** Identify the missing number.
- **Example:** If `nums = [3, 0, 1]`, then `n = 3`, and numbers should be `[0, 1, 2, 3]`. The missing number is `2`.

## **Examples**

### **Example 1**

- **Input:** `nums = [3, 0, 1]`
- **Output:** `2`
- **Explanation:**
  - `n = 3` (length of `nums`)
  - Expected numbers: `[0, 1, 2, 3]`
  - Missing number: `2`

### **Example 2**

- **Input:** `nums = [0, 1]`
- **Output:** `2`
- **Explanation:**
  - `n = 2`
  - Expected numbers: `[0, 1, 2]`
  - Missing number: `2`

### **Example 3**

- **Input:** `nums = [9,6,4,2,3,5,7,0,1]`
- **Output:** `8`
- **Explanation:**
  - `n = 9`
  - Expected numbers: `[0, 1, 2, 3, 4, 5, 6, 7, 8, 9]`
  - Missing number: `8`

## **How to Solve the Problem**

### **Approach 1: Sum Formula**

- **Idea:** The sum of numbers from `0` to `n` can be calculated using the formula for the sum of an arithmetic series.
- **Formula:** `total_sum = n * (n + 1) // 2`
- **Process:**
  1. Calculate the expected total sum of numbers from `0` to `n`.
  2. Calculate the actual sum of elements in `nums`.
  3. The missing number is `total_sum - actual_sum`.

### **Example Using Approach 1**

- **Given `nums = [3, 0, 1]`:**
  - `n = 3`
  - `total_sum = 3 * (3 + 1) // 2 = 6`
  - `actual_sum = 3 + 0 + 1 = 4`
  - **Missing Number:** `6 - 4 = 2`

### **Approach 2: Bit Manipulation (XOR)**

- **Idea:** XORing all numbers from `0` to `n` with all elements in `nums` will cancel out all numbers except the missing one.
- **Properties of XOR:**
  - `a ^ a = 0`
  - `a ^ 0 = a`
  - XOR is associative and commutative.
- **Process:**
  1. Initialize a variable `missing = n`.
  2. For each index `i` from `0` to `n - 1`:
     - `missing ^= i ^ nums[i]`
  3. The result in `missing` is the missing number.

### **Example Using Approach 2**

- **Given `nums = [3, 0, 1]`:**
  - `n = 3`, `missing = 3`
  - Loop through indices and elements:
    - For `i = 0`: `missing ^= 0 ^ nums[0]` ⇒ `missing ^= 0 ^ 3` ⇒ `missing = 3 ^ 0 ^ 3 = 3 ^ 3 = 0`
    - For `i = 1`: `missing ^= 1 ^ nums[1]` ⇒ `missing ^= 1 ^ 0` ⇒ `missing = 0 ^ 1 ^ 0 = 1`
    - For `i = 2`: `missing ^= 2 ^ nums[2]` ⇒ `missing ^= 2 ^ 1` ⇒ `missing = 1 ^ 2 ^ 1 = 2`
  - **Missing Number:** `2`

## **Algorithm Steps**

### **Using Sum Formula**

1. **Calculate `n`:**
   - `n = len(nums)`
2. **Compute Expected Sum:**
   - `total_sum = n * (n + 1) // 2`
3. **Compute Actual Sum:**
   - `actual_sum = sum(nums)`
4. **Find Missing Number:**
   - `missing_number = total_sum - actual_sum`
5. **Return `missing_number`**

### **Using Bit Manipulation**

1. **Initialize `missing`:**
   - `missing = n`
2. **Loop Through Indices:**
   - For `i` from `0` to `n - 1`:
     - `missing ^= i ^ nums[i]`
3. **Return `missing`**

## **Code Examples**

### **Approach 1: Sum Formula (Python)**

```python
def missingNumber(nums):
    n = len(nums)
    total_sum = n * (n + 1) // 2  # Expected sum from 0 to n
    actual_sum = sum(nums)        # Sum of elements in nums
    missing_number = total_sum - actual_sum
    return missing_number
```

### **Approach 2: Bit Manipulation (Python)**

```python
def missingNumber(nums):
    n = len(nums)
    missing = n
    for i in range(n):
        missing ^= i ^ nums[i]
    return missing
```

## **Key Points to Remember**

- **Constraints Handling:**
  - The array contains numbers from `0` to `n` with one missing, and all numbers are unique.
- **Time Complexity:**
  - Both approaches have O(n) time complexity.
- **Space Complexity:**
  - Both approaches use O(1) extra space.
- **Sum Formula Limitations:**
  - Be cautious with potential integer overflow in other languages.
- **Bit Manipulation Benefits:**
  - Avoids issues with integer overflow.
  - Efficient and uses the properties of XOR.

## **Why Learn This?**

- **Algorithmic Thinking:**
  - Enhances problem-solving skills by applying mathematical formulas and bit manipulation.
- **Bit Manipulation:**
  - Understanding XOR operations is valuable for various programming problems.
- **Optimization:**
  - Learn to write algorithms with O(n) time and O(1) space complexity.
- **Technical Interviews:**
  - Commonly asked problem to assess understanding of arrays, math, and bit manipulation.

## **Alternative Approaches**

### **Approach 3: Using a HashSet**

- **Idea:** Store all numbers in a set for O(1) lookups, then check for the missing number.
- **Process:**
  1. Convert `nums` to a set: `num_set = set(nums)`
  2. Loop from `0` to `n`:
     - If a number is not in `num_set`, return it.
- **Complexity:**
  - Time: O(n)
  - Space: O(n) (Not O(1) space)

**Note:** This approach does not meet the O(1) space requirement.

## **Practice Problems**

- **Single Number (LeetCode 136)**
- **Find the Duplicate Number (LeetCode 287)**
- **First Missing Positive (LeetCode 41)**

---

**Remember:** Understanding different approaches helps in selecting the most efficient solution under given constraints. The sum formula and bit manipulation techniques are powerful tools for solving similar problems with optimal space and time complexity.