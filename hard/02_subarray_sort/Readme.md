# Subarray Sort

Write a function that takes in an array of at least two integers and that returns an array of the starting and ending indices of the smallest subarray in the input array that needs to be sorted in place in order for the entire input array to be sorted (in ascending order).

If the input array is already sorted, the function should return [-1, -1].

## Sample Input

```
array = [1, 2, 4, 7, 10, 11, 7, 12, 6, 7, 16, 18, 19]
```

## Sample Output

```
[3, 9]
```

## Hints

```
Hint 1
Realize that even a single out-of-order number in the input array can call for a large subarray to have to be sorted. This is because, depending on how out-of-place the number is, it might need to be moved very far away from its original position in order to be in its sorted position.
```

```
Hint 2
Find the smallest and largest numbers that are out of order in the input array. You should be able to do this in a single pass through the array.
```

```
Hint 3
Once you've found the smallest and largest out-of-order numbers mentioned in Hint #2, find their final sorted positions in the array. This should give you the extremities of the smallest subarray that needs to be sorted.
```

```
Optimal Space & Time Complexity
O(n) time | O(1) space - where n is the length of the input array
```
