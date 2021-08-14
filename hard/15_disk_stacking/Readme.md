# Disk Stacking

You're given a non-empty array of arrays where each subarray holds three integers and represents a disk. These integers denote each disk's width, depth, and height, respectively. Your goal is to stack up the disks and to maximize the total height of the stack. A disk must have a strictly smaller width, depth, and height than any other disk below it.

Write a function that returns an array of the disks in the final stack, starting with the top disk and ending with the bottom disk. Note that you can't rotate disks; in other words, the integers in each subarray must represent [width, depth, height] at all times.

You can assume that there will only be one stack with the greatest total height.

## Sample Input

```
disks = [[2, 1, 2], [3, 2, 3], [2, 2, 8], [2, 3, 4], [1, 3, 1], [4, 4, 5]]
```

## Sample Output
```
[[2, 1, 2], [3, 2, 3], [4, 4, 5]]
// 10 (2 + 3 + 5) is the tallest height we can get by
// stacking disks following the rules laid out above.
```

### Hints

```
Hint 1
Try building an array of the same length as the array of disks. At each index i in this new array, store the height of the tallest tower that can be created with the disk located at index i at the bottom.
```

```
Hint 2
Consider sorting the disks by width, depth, or height for a slight optimization.
```

```
Hint 3
Can you efficiently keep track of potential towers in another array? Instead of storing entire sequences of disks, try storing the indices of previous disks. For example, at index 3 in this other array, store the index of the before-last disk in the tallest tower whose base is the disk at index 3.
```

```
Optimal Space & Time Complexity
O(n^2) time | O(n) space - where n is the number of disks
```
