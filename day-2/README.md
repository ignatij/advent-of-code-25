# Day 2

## Part 1 – brute-force range scan

Each input line is a comma-delimited list of inclusive ID ranges like `11-22`. A range contributes the sum of every ID inside it that consists of some digit sequence repeated exactly twice (e.g., `123123`). The brute-force strategy works fine:

1. Parse the line, split on commas, and for each `lo-hi` iterate every ID `i` in `[lo, hi]`.
2. Convert `i` to decimal and reject immediately if its length is odd. Otherwise compare the first half with the second; add `i` to the running sum when both halves match.

Let `d` be the number of digits per ID (≤ 10) and `N` the total number of IDs across all ranges. The runtime is `O(N * d)` and memory usage stays `O(1)`.

## Part 2 – divisor-based repeated-pattern check

Now an ID is invalid if it is composed of a smaller digit block repeated at least twice, not necessarily exactly two times. The outer enumeration of ranges and IDs is identical; only the string check changes:

1. Convert the candidate ID to decimal text.
2. For every possible block length `k` from 1 up to `len(id)/2`, skip any `k` that doesn’t evenly divide the length.
3. Compare the first block against each subsequent `k`-sized substring; if all match, count the ID and stop checking further `k`.

Because each ID tries all divisors of its length, the complexity is `O(N * d * τ(d))` where `τ(d)` is the number of divisors of the digit count (still tiny in practice). Memory remains `O(1)`.
