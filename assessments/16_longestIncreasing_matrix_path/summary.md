# Summary

We can arrive at the optimal solution to this question by walking through the problem step by step.

In order to find the longest increasing path in the matrix, we're going to have to compare multiple paths in the matrix.

Specifically, we'll want to compare the longest increasing paths that start at each cell in the matrix.

In order to find the longest increasing paths starting at each cell in the matrix (or in order to compute their lengths), we'll have to traverse the matrix multiple times, starting at each cell.

Each traversal should be fairly straightforward; we can recursively traverse through adjacent cells with increasing integers and pick the length of the branch that gives us the greatest path.

But all of these traversals will quickly become redundant, because many of them will rely on previously-done traversals. In other words, we'll be traversing through matrix paths and performing the same expensive time-operations multiple times as we find all of the longest increasing paths.

So the clear solution to avoid doing all of these duplicate computations is to cache the length of every longest increasing path starting at each cell in the matrix. Whenever we visit a cell in a particular traversal, we check our cache to see if we've already computed the length of the longest increasing path starting at that cell.

Implementing this algorithm is fairly straightforward; the code should be self-explanatory.

## Complexity Analysis

The time complexity of this algorithm will be linear with respect to the total number of integers in the matrix; this is because we're only going to do a time-expensive traversal through the entire matrix once, thanks to our caching.

The space complexity of the algorithm will naturally also be linear, since our cache will be the size of the input matrix.

## Closing Thoughts

This is a classic graph-traversal-like problem that can be elegantly and optimally solved using recursion and caching.