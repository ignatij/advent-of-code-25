# Day 11

# Part 1

The input describes a directed graph where each line looks like `you: aaa bbb ccc`. The goal is to count how many distinct routes starting at `you` eventually reach `out`, regardless of length. The solution is to build the graph and run a depth-first search:

- Recurse through every neighbor of the current node.
- Return `1` when a recursive call hits `"out"` to signal that the current branch is a valid path.
- Sum the returned counts from every neighbor to get the total number of paths from the current node.

The graph is small enough that pure recursion already finishes, but I memoized `node -> number of paths` to avoid recomputing identical suffixes. The cache is keyed only by the node name because the state is stateless in this part.

# Part 2

The second puzzle keeps the same graph but only counts paths that touch both `"dac"` and `"fft"` before exiting. I managed this part by including two booleans (`seenDac`, `seenFft`) through the recursion:

- Whenever recursion visits `"dac"` or `"fft"`, flip the corresponding flag to `true`.
- When a branch reaches `"out"`, return `1` only if both flags are true; otherwise return `0` because that route missed at least one checkpoint.
- Continue summing over neighbors exactly like part 1.

Memoization is crucial here because without memoization it never finishes. I memoized `node -> (seenDac, seenFft) -> number of paths` to avoid recomputing identical suffixes.
