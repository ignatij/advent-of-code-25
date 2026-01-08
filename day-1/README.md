# Day 1

Okay, Day 1 of AoC 2025!

## Part 1 – modular dial simulation

Input lists dial rotations (`Lk` or `Rk`) on a 0–99 dial that starts at 50 and wraps around. The password counts how many rotations end with the dial at 0.

1. Parse each instruction, apply the signed distance modulo 100 to update the dial position.
2. After every rotation, increment the counter if the dial now reads 0.

This is O(n) time and O(1) memory—only the current position and total count are tracked.

## Part 2 – zero-crossing arithmetic

Now you count every time the dial clicks past 0, even mid-rotation. A single long rotation can add multiple hits.

1. For `Lk`, measure how many steps remain before hitting 0 when moving left; if the rotation reaches that threshold, count one hit plus extra hits for each additional 100-click lap. Update the dial position modulo 100.
2. Apply the symmetric distance-to-zero logic for `Rk`.

The arithmetic keeps complexity at O(n) time and O(1) memory, avoiding per-click simulation.
