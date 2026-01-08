# Day 6

## Part 1 – columnar parsing left-to-right

The worksheet is a grid where each column belongs to one problem; problems are separated by blank columns. For the standard left-to-right reading order, the algorithm is:

1. Normalize each input row by collapsing every run of spaces into a single space so column boundaries align, then split on spaces to get the digits (or operator) for that column.
2. Iterate column by column. For a given column, read all digits above the bottom row, convert each to an integer, and combine them by either `+` or `*` depending on the operator cell in the bottom row.
3. Sum every column’s result to obtain the grand total.

With `R` rows (including the operator row) and `C` visible columns, parsing is `O(R*C)` time with `O(R*C)` memory for the normalized grid.

## Part 2 – columnar parsing right-to-left

Now the worksheet is interpreted one column at a time from right to left. Digits within a column form numbers vertically (most significant digit at the top), and blank columns still split problems. The solver keeps the raw characters to preserve exact column widths:

1. Scan columns from right to left, accumulating digits for the current problem in a list. When a column is entirely spaces, finalize the current problem by applying the stored operator to the collected numbers (in the order encountered) and add the result to the total.
2. Detect operators from the bottom row as before, updating the current operator whenever the bottom cell contains `+` or `*`.
3. After the loop, reduce the final problem and add it to the total.

This column-wise sweep still touches each cell once, so it remains `O(R*C)` time with `O(R)` auxiliary space for the column accumulator.
