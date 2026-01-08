# Day 8

## Part 1 – Kruskal-style nearest neighbor unions

We’re given 3D coordinates for every junction box (node) and need to connect the 1000 closest pairs by straight-line distance, then multiply the sizes of the three largest resulting connected components. This is essentially the first 1000 edges of Kruskal’s MST process:

1. Parse the `N` coordinates, build all `N*(N-1)/2` edges with squared Euclidean distance, and sort them ascending.
2. Maintain a Union-Find (DSU) structure. For the first 1000 edges in sorted order, union their endpoints.
3. After unions, compute component sizes via another pass over DSU roots, sort sizes descending, and multiply the top three.

Time is dominated by building/sorting all edges: `O(N^2 log N)` (feasible for the input size). DSU operations are near-constant amortized, so total extra time is `O(N^2 α(N))`. Memory is `O(N^2)` for edge storage, plus `O(N)` for DSU.

## Part 2 – last edge to connect all components

Continue the Kruskal process until every node belongs to one component. The last edge that causes the components to unify is the answer; multiply the X-coordinates of its endpoints.

1. Re-use the sorted edge list and DSU from part one (or rebuild). Iterate edges in order, unioning endpoints.
2. After each union, check whether all nodes share one root. Once true, use the current edge’s endpoints to compute `X_A * X_B`.

Runtime remains `O(N^2 log N)` due to edge sorting; each connectedness check is `O(N α(N))` worst case, though you can track component counts to avoid scanning every time. Memory stays `O(N^2)` for edges and `O(N)` for DSU.
