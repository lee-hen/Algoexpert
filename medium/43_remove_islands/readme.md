# Remove Islands

You're given a two-dimensional array (a matrix) of potentially unequal height and width containing only 0s and 1s. The matrix represents a two-toned image, where each 1 represents black and each 0 represents white. An island is defined as any number of 1s that are horizontally or vertically adjacent (but not diagonally adjacent) and that don't touch the border of the image. In other words, a group of horizontally or vertically adjacent 1s isn't an island if any of those 1s are in the first row, last row, first column, or last column of the input matrix.

Note that an island can twist. In other words, it doesn't have to be a straight vertical line or a straight horizontal line; it can be L-shaped, for example.

You can think of islands as patches of black that don't touch the border of the two-toned image.

Write a function that returns a modified version of the input matrix, where all of the islands are removed. You remove an island by replacing it with 0s.

Naturally, you're allowed to mutate the input matrix.

## Sample Input
```
matrix = 
[
  [1, 0, 0, 0, 0, 0],
  [0, 1, 0, 1, 1, 1],
  [0, 0, 1, 0, 1, 0],
  [1, 1, 0, 0, 1, 0],
  [1, 0, 1, 1, 0, 0],
  [1, 0, 0, 0, 0, 1],
]
```

## Sample Output
```
[
  [1, 0, 0, 0, 0, 0],
  [0, 0, 0, 1, 1, 1],
  [0, 0, 0, 0, 1, 0],
  [1, 1, 0, 0, 1, 0],
  [1, 0, 0, 0, 0, 0],
  [1, 0, 0, 0, 0, 1],
] 
// The islands that were removed can be clearly seen here:
// [
//   [ ,  ,  ,  ,  , ],
//   [ , 1,  ,  ,  , ],
//   [ ,  , 1,  ,  , ],
//   [ ,  ,  ,  ,  , ],
//   [ ,  , 1, 1,  , ],
//   [ ,  ,  ,  ,  , ], 
// ]
```

### Hints

```
Hint 1
How would you solve this problem if you knew the positions of all the non-island 1s?
```

```
Hint 2
Find and store the positions of all the non-island 1s. You can do this by performing a graph traversal (depth-first search, for example) on all the 1s that are on the border of the image. Afterwards, you can easily identify and remove all the island 1s from the input matrix by relying on the data structure that you used to store the positions of non-island 1s.
```

```
Hint 3
You can also solve this problem without the use of a data structure that stores the positions of non-islands 1s. Simply loop through the border of the image, and perform a depth-first search on all positions with the value 1. During this depth-first search, find all the 1s that are connected to the original position on the border, and change them from 1 to 2. After changing all non-island 1s to 2s, you can simply remove all the remaining 1s, which are guaranteed to be islands, from the matrix (by replacing them with 0s), and you can then change all the 2s back to 1s, since these were previously determined to be non-islands.
```

```
Optimal Space & Time Complexity
O(wh) time | O(wh) space - where w and h are the width and height of the input matrix
```
