# Pascal's Triangle

## **Problem Summary**

- **Goal:** Generate the first `numRows` of Pascal's Triangle.
- **Input:** An integer `numRows` (`1 ≤ numRows ≤ 30`).
- **Output:** A list of lists, where each inner list represents a row of Pascal's Triangle.

## **Understanding Pascal's Triangle**

- **Structure:** A triangular array of numbers.
- **Rule:** Every number is the sum of the two numbers directly above it.
- **First Row:** Starts with a single `1`.

### **Visual Example (First 5 Rows)**

```
Row 1:        [1]
Row 2:       [1, 1]
Row 3:      [1, 2, 1]
Row 4:     [1, 3, 3, 1]
Row 5:    [1, 4, 6, 4, 1]
```

## **How to Build Pascal's Triangle**

1. **Start with the first row:** `[1]`.
2. **For each subsequent row:**
   - **Begin and end with `1`.**
   - **Calculate middle elements:**
     - Each middle number is the sum of the two numbers directly above it from the previous row.

### **Example: Building the 4th Row**

- **Previous Row (3rd Row):** `[1, 2, 1]`
- **New Row Start:** `[1]`
- **Calculate Middle Elements:**
  - `1 + 2 = 3`
  - `2 + 1 = 3`
- **New Row So Far:** `[1, 3, 3]`
- **Add Ending `1`:**
- **Completed Row:** `[1, 3, 3, 1]`

## **Algorithm Steps**

1. **Initialize** an empty list `result` to hold all rows.
2. **Loop** from `i = 0` to `numRows - 1`:
   - **Create** a new row with `i + 1` elements, all initialized to `1`.
   - **Update** middle elements for `j` from `1` to `i - 1`:
     - `row[j] = result[i - 1][j - 1] + result[i - 1][j]`
   - **Append** the new row to `result`.
3. **Return** `result`.

## **Simple Code Example (Python)**

```python
def generate(numRows):
    result = []
    for i in range(numRows):
        # Start with a row of '1's
        row = [1] * (i + 1)
        # Compute the middle elements
        for j in range(1, i):
            row[j] = result[i - 1][j - 1] + result[i - 1][j]
        # Add the row to the result
        result.append(row)
    return result
```

## **Key Points to Remember**

- **Edges are always `1`**: The first and last elements of each row are `1`.
- **Middle elements**: Calculated by summing the two numbers above.
- **Building process**: Each row is built using the previous row.
- **Constraints**: Ensure your solution works efficiently for `1 ≤ numRows ≤ 30`.

## **Why Learn This?**

- **Foundation in Combinatorics**: Pascal's Triangle relates to binomial coefficients.
- **Algorithm Practice**: Good exercise for nested loops and array manipulation.
- **Problem-Solving Skills**: Enhances understanding of how to build structures iteratively.

---

**Remember:** Start simple, focus on understanding how each row relates to the previous one, and build up the triangle step by step!