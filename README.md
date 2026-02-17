# Nethermind Technical Skills Assessment

Technical skills assessment created by Nethermind through Codility, solved in Go.

## Problems

### P1 - Maximize Car Production

A car manufacturer wants to maximize the number of cars produced in a month.
Given an array `H` where `H[K]` is the hours needed to produce car `K`,
and two assembly lines with `X` and `Y` working hours respectively:

- Each car is produced on exactly one line (no switching mid-production).
- Only one car at a time per line.
- Find the **maximum number of different cars** that can be produced.

**Approach:** Sort cars by production time (ascending), then use incremental 0-1 knapsack DP
to check if the `k` cheapest cars can be partitioned across the two lines.

**Complexity:** O(N \* X) time, O(X) space.

### P2 - Longest Alternating Path in a Tree

Given a tree of `N` nodes (each labeled `'a'` or `'b'`), represented by a parent array `A`
and a string `S`:

- `A[K]` is the parent of node `K` (`A[0] = -1` for the root).
- `S[K]` is the letter at node `K`.
- Find the **number of vertices on the longest path** where no two adjacent vertices share the same letter.

**Approach:** BFS from root to establish processing order, then bottom-up DP.
For each node, track the longest alternating downward path. Combine the two best
child branches through each node to find the global longest path.

**Complexity:** O(N) time, O(N) space.

## Running

```bash
# Run tests for a specific problem
go test ./problems/p1/
go test ./problems/p2/

# Run all tests
go test ./...
```
