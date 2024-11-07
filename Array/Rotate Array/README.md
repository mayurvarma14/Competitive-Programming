# Rotate Array

## **Problem Summary**

- **Goal:** Rotate an array `nums` to the right by `k` steps.
- **Input:**
  - An integer array `nums`.
  - A non-negative integer `k`.
- **Output:** Modify `nums` in-place to reflect the rotated array.

**Constraints:**

- `1 ≤ nums.length ≤ 10⁵`
- `-2³¹ ≤ nums[i] ≤ 2³¹ - 1`
- `0 ≤ k ≤ 10⁵`

## **Understanding the Problem**

- **Array Rotation:** Moving elements in the array so that each element shifts to the right by `k` positions.
- **Wrapping Around:** Elements shifted beyond the array's end wrap around to the beginning.
- **In-Place Requirement:** The rotation should be done without allocating extra space for another array (O(1) space complexity).

### **Example 1**

- **Input:** `nums = [1, 2, 3, 4, 5, 6, 7]`, `k = 3`
- **Output:** `[5, 6, 7, 1, 2, 3, 4]`
- **Explanation:**
  - Rotate 1 step: `[7, 1, 2, 3, 4, 5, 6]`
  - Rotate 2 steps: `[6, 7, 1, 2, 3, 4, 5]`
  - Rotate 3 steps: `[5, 6, 7, 1, 2, 3, 4]`

### **Example 2**

- **Input:** `nums = [-1, -100, 3, 99]`, `k = 2`
- **Output:** `[3, 99, -1, -100]`
- **Explanation:**
  - Rotate 1 step: `[99, -1, -100, 3]`
  - Rotate 2 steps: `[3, 99, -1, -100]`

## **How to Rotate the Array**

### **Method 1: Using Extra Space (Not In-Place)**

1. **Create a New Array:**
   - Allocate a new array of the same length as `nums`.
2. **Copy Elements:**
   - For each index `i`, place `nums[i]` at the position `(i + k) % n` in the new array.
3. **Copy Back:**
   - Copy the new array back into `nums`.

- **Space Complexity:** O(n) (Not acceptable if O(1) space is required).

### **Method 2: Reverse Parts of the Array (In-Place)**

**Idea:** Reverse the entire array and then reverse parts of it to achieve the rotation.

1. **Calculate Effective Rotations:**
   - Since rotating `n` times results in the original array, use `k = k % n`.

2. **Reverse Entire Array:**
   - Reverse all elements in `nums`.

3. **Reverse First `k` Elements:**
   - Reverse elements from index `0` to `k - 1`.

4. **Reverse Remaining Elements:**
   - Reverse elements from index `k` to `n - 1`.

**Why This Works:**

- **After Step 2:** The array is completely reversed.
- **After Step 3:** The first `k` elements are in the correct rotated position.
- **After Step 4:** The remaining elements are also in the correct position.

### **Visual Example**

**Given `nums = [1, 2, 3, 4, 5, 6, 7]`, `k = 3`:**

1. **Original Array:** `[1, 2, 3, 4, 5, 6, 7]`
2. **Reverse Entire Array:** `[7, 6, 5, 4, 3, 2, 1]`
3. **Reverse First `k=3` Elements:** `[5, 6, 7, 4, 3, 2, 1]`
4. **Reverse Remaining Elements:** `[5, 6, 7, 1, 2, 3, 4]`

## **Algorithm Steps**

1. **Calculate Array Length:**
   - `n = len(nums)`

2. **Adjust `k`:**
   - `k = k % n` (Handles cases where `k` ≥ `n`)

3. **Define Reverse Function:**
   - A helper function to reverse elements between two indices.

4. **Reverse Entire Array:**
   - `reverse(nums, 0, n - 1)`

5. **Reverse First `k` Elements:**
   - `reverse(nums, 0, k - 1)`

6. **Reverse Remaining Elements:**
   - `reverse(nums, k, n - 1)`

## **Intuitive Code Example (Python)**

```python
def rotate(nums, k):
    n = len(nums)
    k = k % n  # Adjust k if it's greater than n

    def reverse(arr, start, end):
        while start < end:
            arr[start], arr[end] = arr[end], arr[start]  # Swap elements
            start += 1
            end -= 1

    # Reverse the entire array
    reverse(nums, 0, n - 1)
    # Reverse the first k elements
    reverse(nums, 0, k - 1)
    # Reverse the remaining elements
    reverse(nums, k, n - 1)
```

**Explanation:**

- **`reverse` Function:** Swaps elements from both ends moving towards the center.
- **In-Place:** Modifies `nums` directly without extra space.

## **Key Points to Remember**

- **In-Place Rotation:** Use reversing to achieve O(1) space complexity.
- **Adjust `k`:** Always compute `k % n` to handle large `k` values.
- **Reverse Operations:**
  - Reversing sections of the array rearranges elements efficiently.
- **Helper Function:** Using a `reverse` function keeps the code clean and modular.

## **Alternative Methods**

### **Method 3: Cyclic Replacements**

1. **Calculate Effective Rotations:**
   - `k = k % n`

2. **Initialize Count of Swapped Elements:**
   - `count = 0`

3. **Loop Through Positions:**
   - For each starting index `i` from `0` to `gcd(n, k) - 1`:
     - Set `current = i`
     - Store `prev_value = nums[i]`
     - While `True`:
       - Compute `next_idx = (current + k) % n`
       - Swap `nums[next_idx]` with `prev_value`
       - Update `current = next_idx`
       - Increment `count`
       - Break when `current == i`
     - If `count >= n`, break

- **Complexity:**
  - **Time:** O(n)
  - **Space:** O(1)

### **Method 4: Using Slicing (Not In-Place)**

- **Note:** This method is not in-place and uses extra space.

```python
def rotate(nums, k):
    n = len(nums)
    k = k % n
    nums[:] = nums[-k:] + nums[:-k]
```

- **Explanation:** Reassigns `nums` to a new list combining the last `k` elements and the first `n - k` elements.

## **Why Learn This?**

- **Array Manipulation:** Understanding how to manipulate arrays efficiently.
- **Algorithm Optimization:** Learning in-place algorithms that optimize space complexity.
- **Problem-Solving Skills:** Enhances ability to approach problems with different strategies.
- **Technical Interviews:** Commonly asked question to test understanding of array operations and algorithmic thinking.

## **Practice Problems**

- **Rotate Matrix (Image Rotation)**
- **Reverse Words in a String**
- **Shift 2D Grid**

---

**Remember:** Rotating an array in-place can be efficiently achieved by reversing parts of the array. Always consider edge cases like when `k` is larger than the array length. Understanding different methods enhances problem-solving flexibility!