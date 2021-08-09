# Longest Increasing Matrix Path

Write a function that takes in a non-empty two-dimensional integer array (a matrix) of potentially unequal height and width and returns the length of its longest increasing path.

A matrix path is a collection of matrix elements where each element is horizontally or vertically adjacent to the next element.

A matrix path is increasing if each of its elements is strictly greater than the previous element.

## Sample Input

``` 
matrix = [
  [1,  2,  4,  3,  2],
  [7,  6,  5,  5,  1],
  [8,  9,  7, 15, 14],
  [5, 10, 11, 12, 13],
]
```

## Sample Output

``` 
15

// The longest increasing path can be clearly
// seen here, starting at 1 and ending at 15:
// [
//   [ ,   ,  4,  3,  2],
//   [7,  6,  5,   ,  1],
//   [8,  9,   , 15, 14],
//   [ , 10, 11, 12, 13],
// ]
```
