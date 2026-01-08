# Day 3

## Part 1 – brute force two-digit scan

Each line is a “bank” of digit-labeled batteries. Turning on exactly two batteries produces a two-digit number (order fixed by their positions), and we want the maximum per bank.

1. For a bank string of length `m`, consider every ordered pair `(i, j)` with `i < j`, form the two-digit value using digits at those positions, and track the largest.
2. Sum those per-bank maxima.

With `m ≤ 15`, brute force costs `O(m^2)` per bank; streaming the input keeps memory at `O(1)` aside from the current line.

## Part 2 – monotonic stack for largest subsequence

Now we must choose exactly twelve digits per bank to maximize the resulting 12-digit number (relative order preserved). This is equivalent to removing `k = len(bank) - 12` digits while keeping the sequence lexicographically largest.

1. Traverse the digits left to right while maintaining a stack. Whenever the current digit is greater than the stack’s top and we still have deletions left (k > 0), pop from the stack; this greedily discards smaller prefixes.
2. After processing, the stack may still be longer than 12 (if we never spent all deletions on the fly), so truncate the tail until exactly 12 digits remain.
3. Parse the resulting digits as an integer and accumulate across banks.

This is the classic “build the largest subsequence of fixed length” algorithm, running in `O(m)` per bank with `O(m)` extra space for the stack.
