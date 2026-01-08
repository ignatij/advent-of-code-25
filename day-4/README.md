# Day 4

## Part 1 - neighbor-count grid scan

The input is a grid of `.` and `@` where each `@` represents a paper roll. A roll is accessible if fewer than four of its eight neighbors (N, NE, E, SE, S, SW, W, NW) are also `@`.

1. Parse the grid into a 2D array of runes or bytes.
2. For every `@`, count occupied neighbors by scanning the 8 offsets and stay within bounds.
3. Increment the answer whenever the neighbor count is below 4.

Runtime is `O(R*C)` where R is the number of rows and C is the number of columns because each cell is visited once and performs constant work. Memory is `O(R*C)` for the stored grid (or `O(1)` if processed streaming with a buffer).

## Part 2 - iterative neighbor-pruning

Now rolls get removed iteratively: every pass removes all currently accessible rolls, which may unlock new rolls for the next pass. Keep repeating until no roll qualifies.

1. Loop over the grid and mark every `@` with <4 neighbors; immediately turn them into `.` and add to the running total.
2. Track whether at least one removal happened during the pass; if not, stop.
3. Repeat the scan-removal cycle while removals occur.

Runtime remains `O(R*C)` per pass (each pass scans every cell once), but in the worst case there can be `O(R*C)` passes—removing one roll per iteration—so the total worst-case time is `O((R*C)^2)`. Memory stays `O(R*C)` for the grid.
