# Summary

We can solve this question in three different ways, each more optimal than the previous.

## Least Optimal Solution

We can iterate through the array once, and at each index, iterate through the left side and right side of the current index to compute the leftSideSum and rightSideSum.

If the two sums are ever equal to each other, we've found the balance index.

This solution has us iterate through the entire array at every index of the array; therefore, it runs in O(n^2) time, where n is the length of the array.

## More Optimal Solution

Instead of using nested for loops to calculate the leftSideSums and rightSideSums, we can compute them trivially by iterating through the array twice: once from left to right for the leftSideSums, and once from right to left for the rightSideSums.

By initializing the two sums to 0, all that we have to do is to increment the sums at each index with the current number in the array; this gives us the relevant sum for the next index.

While this solution lets us avoid the nested loops, which improves the time complexity of our algorithm to linear time, it does require us to store at least the leftSideSums in an auxiliary array, forcing us to use linear space.

## Most Optimal Solution

For the most optimal solution, we have to realize that both the leftSideSums and the rightSideSums can be computed in a single pass through the array, without needing to store an entire auxiliary array of values.

To do this, we initialize the leftSideSum to 0 and the rightSideSum to the sum of the values in the array, which is precomputed trivially with one pass through the array.

Then, we iterate through the array, and at each index we decrement the rightSideSum and increment the leftSideSum by the value of the current number. We also check for equality of the two sums, and we make sure to sequence these three operations correctly.

Since we only do a total of two passes through the array and store just a few variables, this solution runs in linear time and with constant space.

## Complexity Analysis

As detailed above, the optimal solution to this problem runs in linear time and with constant space.

## Closing Thoughts

One of the easiest ways to solve this problem optimally is to start with the naive solution and to work your way through optimizations.
