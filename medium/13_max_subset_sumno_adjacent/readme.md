# Max Subset Sum No Adjacent

Write a function that takes in an array of positive integers and returns the maximum sum of non-adjacent elements in the array.

If the input array is empty, the function should return 0.

## Sample Input

```
array = [75, 105, 120, 75, 90, 135]
```

## Sample Output

```
330 // 75 + 120 + 135
```

### Hints

```
Hint 1
Try building an array of the same length as the input array. At each index in this new array, store the maximum sum that can be generated using no adjacent numbers located between index 0 and the current index.
```

```
Hint 2
Can you come up with a formula that relates the max sum at index i to the max sums at indices i - 1 and i - 2?
```

```
Hint 3
Do you really need to store the entire array mentioned in Hint #1, or can you somehow store just the max sums that you need at any point in time?
```

```
Optimal Space & Time Complexity
O(n) time | O(1) space - where n is the length of the input array
```
