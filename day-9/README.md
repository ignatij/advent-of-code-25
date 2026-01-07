# Day 9

# Part 1

Every input line is a lattice point that lies on the outline of some axis-aligned rectangle. Because the file is tiny I simply check every pair of points: two corners determine a rectangle whose side lengths are `|x2-x1|+1` by `|y2-y1|+1`, so its area is `(abs(dx)+1)*(abs(dy)+1)`. Track the maximum over all `O(n^2)` pairs and print it. Brute force is fine here and keeps the code trivial.

Eazy.

# Part 2

Boy oh boy, I was in for a surprise.

Now we must ignore any rectangle that intersects the walkways traced by the input polygon. Tried with a naive BFS over the entire coordinate range, it blew up because the coordinates can be huge and sparse.

The fix has three ingredients:

1. **Coordinate compression.** Collect every x/y coordinate that matters (each point plus ±1 neighbors), sort/deduplicate them, and map real coordinates to compact indices. That lets us build a manageable grid even when raw coordinates are large.
2. **Flood fill of “outside” cells.** Using the compressed grid, mark every unit square that lies outside the polygon by running a BFS from the outer boundary, respecting vertical/horizontal “walls” built from the polygon edges.
3. **2D prefix sums.** After flood fill, build a prefix-sum grid of the outside cells. For any candidate rectangle defined by two original points, convert its bounds to compressed indices and query the prefix matrix to see if any outside cells lie inside that rectangle. If the count is zero, it’s valid, and we compute its area in original coordinates.

This combination keeps the search finite and lets us test each pair of corners in constant time after preprocessing. The pairwise loop is still `O(n^2)`, but now each test simply checks prefix sums rather than attempting another expensive traversal.

This really took a while.
