# Day 5

# Part 1 - interval-merging and binary-search

Input consists of a set of inclusive fresh-ID ranges, a blank line, and then a list of candidate IDs. An ID is fresh if it falls inside any range (ranges may overlap). The approach:

1. Parse all ranges into `(start, end)` pairs, sort by `start`, and merge overlaps so that the range list is disjoint and sorted.
2. For each candidate ID, run binary search over the merged ranges to check membership; if the range at `mid` contains the ID, count it.

Let `R` be the number of original ranges and `F` the number of candidate IDs. Sorting and merging costs `O(R log R)`; each membership test is `O(log R)` so the total is `O(R log R + F log R)` with `O(R)` memory for the merged list.

# Part 2 - interval-merging

The second part ignores the candidate list. After merging the ranges exactly as in part one, simply sum the lengths of the disjoint ranges (`end - start + 1` per range). The result is the total number of IDs flagged as fresh.

Runtime is dominated by the same `O(R log R)` sort/merge step, and memory usage stays `O(R)` for the merged ranges.
