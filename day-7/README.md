# Day 7

## Part 1 – recursive beam split counting

The tachyon manifold is a grid with a single source `S` at the top and splitters `^`. Beams flow only downward; whenever a beam hits a splitter, it stops and spawns new beams immediately to the left and right that continue downward. We only need the number of split events.

1. Parse the grid and locate `S`. Maintain a `visited` matrix to avoid revisiting the same `(row, col)` from above.
2. Recurse downward from `S`. If the current cell is empty (`.`), continue to the next row. If it’s a splitter, add 1 to the count and recursively process the left and right neighbors on the same row (each then continues downward).
3. Stop when leaving the grid or revisiting a cell.

This DFS-style traversal touches each reachable cell at most once, so runtime is `O(R*C)` (where `R` is the number of rows and `C` is the number of columns) and memory is `O(R*C)` for the grid plus visited bitmap.

## Part 2 – dynamic programming over split DAG

Now a single tachyon particle creates separate timelines for every possible left/right choice at splitters. Each splitter contributes the sum of timelines from its two children (down-left and down-right on the next row), and empty cells simply inherit the count from directly below. Bottom-row cells represent terminal timelines.

1. Build a DP table `dp[row][col]` representing the number of timelines reaching that cell.
2. Initialize the bottom row to 1s because once the particle reaches the bottom, it ends the journey.
3. Fill the table bottom-up: if `(row, col)` is a splitter, set `dp[row][col] = dp[row+1][col-1] + dp[row+1][col+1]` (guarding bounds). Otherwise copy `dp[row+1][col]`.
4. The answer is `dp[startRow][startCol]`.

The grid is processed once, giving `O(R*C)` time and `O(R*C)` memory for the DP table.
