# Remove Duplicates from Sorted Array

## **Problem Summary**

- **Goal:** Remove duplicates from a sorted array `nums` **in-place** so that each unique element appears only once.
- **Input:** A sorted integer array `nums` in non-decreasing order.
- **Output:** An integer `k` representing the number of unique elements. Modify `nums` so that the first `k` elements are the unique elements in their original order.

**Constraints:**

- `1 ≤ nums.length ≤ 3 × 10⁴`
- `-100 ≤ nums[i] ≤ 100`
- `nums` is sorted in non-decreasing order.

## **Understanding the Problem**

- **In-Place Modification:** Modify `nums` without using extra space (constant O(1) extra memory).
- **Unique Elements:** After modification, the first `k` elements of `nums` should be the unique elements.
- **Order Matters:** The relative order of the unique elements must be maintained.
- **Extra Elements:** Values beyond the first `k` elements are not important.

### **Custom Judge Explanation**

- **Validation Steps:**
  1. Check that your function returns the correct `k`.
  2. Verify that the first `k` elements of `nums` are the unique elements.
  3. Ensure the elements are in the same order as the original array.

## **Examples**

### **Example 1**

- **Input:** `nums = [1, 1, 2]`
- **Output:** `k = 2`, `nums = [1, 2, _]`
- **Explanation:**
  - Unique elements: `[1, 2]`
  - The function should return `k = 2`.
  - Modify `nums` so that the first two elements are `1` and `2`.
  - The `_` denotes irrelevant values beyond the first `k` elements.

### **Example 2**

- **Input:** `nums = [0, 0, 1, 1, 1, 2, 2, 3, 3, 4]`
- **Output:** `k = 5`, `nums = [0, 1, 2, 3, 4, _, _, _, _, _]`
- **Explanation:**
  - Unique elements: `[0, 1, 2, 3, 4]`
  - The function should return `k = 5`.
  - Modify `nums` so that the first five elements are `0, 1, 2, 3, 4`.

## **How to Remove Duplicates In-Place**

### **Two-Pointer Technique with Meaningful Variable Names**

1. **Initialize Two Pointers:**
   - **`write_index`:** Indicates the position where the next unique element should be written.
   - **`read_index`:** Used to traverse the array and find unique elements.

2. **Process:**
   - Start with `write_index = 0` (assuming the first element is unique).
   - For `read_index` from `1` to `len(nums) - 1`:
     - **Compare:** If `nums[read_index]` is not equal to `nums[write_index]`:
       - **Increment `write_index`:** Move to the next position for a unique element.
       - **Assign:** Set `nums[write_index] = nums[read_index]` to store the new unique element.

3. **Result:**
   - The first `write_index + 1` elements of `nums` are the unique elements.
   - Return `k = write_index + 1` as the number of unique elements.

### **Visual Example**

**Given `nums = [0, 0, 1, 1, 1, 2, 2, 3, 3, 4]`:**

| Step | `write_index` | `read_index` | `nums`                                  | Action                                            |
|------|---------------|--------------|------------------------------------------|---------------------------------------------------|
| Init | 0             | 1            | `[0, 0, 1, 1, 1, 2, 2, 3, 3, 4]`        | Start pointers                                    |
| 1    | 0             | 1            | `nums[1] = 0` equals `nums[0] = 0`      | No action                                         |
| 2    | 0             | 2            | `nums[2] = 1` not equal to `nums[0] = 0`| Increment `write_index` to 1, set `nums[1] = 1`    |
| 3    | 1             | 3            | `nums[3] = 1` equals `nums[1] = 1`      | No action                                         |
| 4    | 1             | 4            | `nums[4] = 1` equals `nums[1] = 1`      | No action                                         |
| 5    | 1             | 5            | `nums[5] = 2` not equal to `nums[1] = 1`| Increment `write_index` to 2, set `nums[2] = 2`    |
| 6    | 2             | 6            | `nums[6] = 2` equals `nums[2] = 2`      | No action                                         |
| 7    | 2             | 7            | `nums[7] = 3` not equal to `nums[2] = 2`| Increment `write_index` to 3, set `nums[3] = 3`    |
| 8    | 3             | 8            | `nums[8] = 3` equals `nums[3] = 3`      | No action                                         |
| 9    | 3             | 9            | `nums[9] = 4` not equal to `nums[3] = 3`| Increment `write_index` to 4, set `nums[4] = 4`    |

- **Final `nums`:** `[0, 1, 2, 3, 4, 2, 2, 3, 3, 4]`
- **Unique Elements:** Positions `0` to `4` (`write_index = 4`)
- **Return `k = write_index + 1 = 5`**

## **Algorithm Steps**

1. **Edge Case Check:**
   - If `nums` is empty, return `0`.

2. **Initialize Pointers:**
   - Set `write_index = 0` (first unique element is at index 0).

3. **Traverse the Array:**
   - For `read_index` from `1` to `len(nums) - 1`:
     - **If `nums[read_index] != nums[write_index]`:**
       - Increment `write_index` by `1`.
       - Assign `nums[write_index] = nums[read_index]`.

4. **Return Result:**
   - Return `k = write_index + 1`.

## **Intuitive Code Example (Python)**

```python
def removeDuplicates(nums):
    if not nums:
        return 0

    write_index = 0  # Index where the next unique element should be written

    # Traverse the array starting from the second element
    for read_index in range(1, len(nums)):
        # If the current element is different from the last unique element
        if nums[read_index] != nums[write_index]:
            write_index += 1  # Move to the next position for a unique element
            nums[write_index] = nums[read_index]  # Update the unique element

    # The number of unique elements is write_index + 1
    return write_index + 1
```

**Explanation:**

- **Variables:**
  - `write_index`: Position where the next unique element should be placed.
  - `read_index`: Position of the current element being read.

- **Logic Flow:**
  - Start with `write_index = 0` because the first element is always unique.
  - Iterate `read_index` from `1` to the end of the array.
  - If `nums[read_index]` is not equal to `nums[write_index]`, increment `write_index` and update `nums[write_index]` with `nums[read_index]`.

## **Key Points to Remember**

- **Two-Pointer Technique:**
  - Efficient for in-place array manipulation without extra space.
  - One pointer (`read_index`) reads through the array.
  - Another pointer (`write_index`) writes unique elements.

- **In-Place Modification:**
  - Do not allocate extra space for another array.
  - Modify the input array directly.

- **Maintain Order:**
  - Keep the original order of unique elements.
  - Only overwrite when a new unique element is found.

- **Return Value:**
  - The number of unique elements is `write_index + 1`.

## **Why Learn This?**

- **Fundamental Algorithm:**
  - The two-pointer technique is commonly used in array problems.

- **Space Efficiency:**
  - Teaches how to optimize for O(1) space complexity.

- **Interview Favorite:**
  - Popular question to assess understanding of pointers and in-place operations.

- **Problem-Solving Skills:**
  - Enhances ability to manipulate data structures efficiently.

---

**Remember:** By using meaningful variable names like `read_index` and `write_index`, the logic becomes clearer, making the code more readable and easier to understand. The two-pointer technique allows you to efficiently remove duplicates while maintaining the relative order of elements in the array.