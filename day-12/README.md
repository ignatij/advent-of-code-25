# Day 12

Well this was a nice surprise!

The input describes a set of oddly shaped presents plus several rectangular regions under Christmas trees. Each region line looks like `12x5: 1 0 1 0 2 2`, meaning the rectangle is 12 units wide by 5 tall and we must pack the indicated number of presents of each shape. Presents can be rotated and flipped, must stay on the integer grid, and can’t overlap, so naively this is a tiling/packing problem that looks NP-hard.

First time I read the problem statement, I thought it would be a mission impossible!
Peeking at the actual puzzle input (and I admin the [AoC Subreddit](https://www.reddit.com/r/adventofcode/)) revealed a huge simplification: every present shape has exactly 9 cells (3×3 area). The only feasibility test needed is checking whether the combined area of the required presents is at most the area of the region.
