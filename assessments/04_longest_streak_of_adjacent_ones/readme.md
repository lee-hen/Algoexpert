# Longest Streak Of Adjacent Ones

Write a function that takes in a non-empty array of only 0s and 1s and returns the index of the 0 that, if replaced by a 1, would form the longest streak of adjacent 1s in the entire array.

If there are no 0s in the array, your function should return -1; if there are multiple possible answers, your function can return any of them.

## Sample Input
```
  array = [1, 0, 1, 1, 0, 0, 1, 1, 0, 1, 1, 1, 0, 1]
```

## Sample Output
``` 
  8 // Replacing the 0 at index 8 with a 1 forms a streak of 6 adjacent 1s.
```