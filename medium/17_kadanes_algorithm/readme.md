# Kadane's Algorithm

Write a function that takes in a non-empty array of integers and returns the maximum sum that can be obtained by summing up all of the integers in a non-empty subarray of the input array. A subarray must only contain adjacent numbers (numbers next to each other in the input array).

## Sample Input

```
array = [3, 5, -9, 1, 3, -2, 3, 4, 7, 2, -9, 6, 3, 1, -5, 4]
```

## Sample Output

```
19 // [1, 3, -2, 3, 4, 7, 2, -9, 6, 3, 1]
```

## Hints

```
Hint 1
This problem seems fairly simple until you run into negative numbers, some of which are so big in absolute value that they essentially break an otherwise good subarray into two subarrays, and some of which are small enough that there exists a subarray containing them whose numbers sum to maximum sum that you're looking for. How can you determine which group a negative number belongs to?
```

```
Hint 2
Realize that at any given index in the input array, the maximum sum for a subarray ending at that index is either the maximum sum for a subarray ending at the previous index plus the number at that index, or just the number at that index. Thus, for each index in the array, you can calculate the maximum sum of a subarray ending at that index, and this can be done in one simple pass through the input array.
```

```
Hint 3
How can you alter the pass through the input array mentioned in Hint #2 so as to obtain the actual answer to the problem, that is the maximum sum of any subarray in the input array? You should be able to accomplish everything in one loop through the input array.
```

```
Optimal Space & Time Complexity
O(n) time | O(1) space - where n is the length of the input array
```
