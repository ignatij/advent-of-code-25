# Day 10

# Part 1 – meet-in-the-middle XOR search

Every indicator row becomes a bitmask: `#` means the corresponding bit must be 1 in the final state, `.` means 0. Each button is also a bitmask with 1s wherever that button toggles. With that encoding the puzzle says: “Pick a multiset of button masks whose XOR equals the target mask while minimizing presses.” Enumerating all subsets is `O(2^n)` and infeasible, so I split the buttons into two halves and used meet-in-the-middle:

1. For the left half build a map `state -> min_presses` containing all XOR states reachable by subsets of that half.
2. Do the same for the right half.
3. For each right-state `r`, the matching left-state must be `target XOR r`; sum the minimal press counts from both tables and keep the overall minimum.

Because each half only has `n/2` buttons the enumeration is fast while still covering the entire space.

# Part 2 – integer linear programming with Z3

Same as for Day 9, for Part 2 things got pretty interesting again.

Here each button adds 1 to multiple counters, so presses are no longer modular arithmetic—they’re integer quantities. I tried a DFS/backtracking search that picks how many times to press each button, but the search space is enormous (targets are in the hundreds), so it never finished.

Instead I recast the machine as an Integer Linear Program (ILP). Give every button `j` an integer variable `x_j >= 0` representing how many times we press it. For each counter `i`, sum the `x_j` of every button touching that counter, multiply by the number of times it touches it, and insist that the sum equals the target joltage `b_i`. That yields linear equality constraints. The objective “fewest total presses” is simply `minimize sum_j x_j`.

I solved this ILP with the Go bindings for Z3 (<https://pkg.go.dev/github.com/aclements/go-z3/z3>). The binding doesn’t expose Z3’s Optimize API, so I binary-search the minimum objective: assert all constraints once, then repeatedly push an extra inequality `sum_j x_j <= mid`, check satisfiability, and adjust the bound until the smallest feasible `mid` is found. Summing those minima over all machines gives the final answer.

This was definetely the hardest challenge of the AoC 2025!
Probably a big contributing fact for that would be my lack of experience with ILP problems.
