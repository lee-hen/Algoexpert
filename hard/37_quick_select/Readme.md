# Quickselect

Write a function that takes in an array of distinct integers as well as an integer k and that returns the kth smallest integer in that array.

The function should do this in linear time, on average.

## Sample Input

```
array = [8, 5, 2, 9, 7, 6, 3]
k = 3
```

## Sample Output
```
5
```

### Hints

Hint 1
> The Quick Sort sorting algorithm works by picking a "pivot" number from an array, positioning every other number in the array in sorted order with respect to the pivot (all smaller numbers to the pivot's left; all bigger numbers to the pivot's right), and then repeating the same two steps on both sides of the pivot until the entire array is sorted. Apply the technique used in Quick Sort until the pivot element gets positioned in the kth place in the array, at which point you'll have found the answer to the problem.

Hint 2
> Pick a random number from the input array (the first number, for instance) and let that number be the pivot. Iterate through the rest of the array using two pointers, one starting at the left extremity of the array and progressively moving to the right, and the other one starting at the right extremity of the array and progressively moving to the left. As you iterate through the array, compare the left and right pointer numbers to the pivot. If the left number is greater than the pivot and the right number is less than the pivot, swap them; this will effectively sort these numbers with respect to the pivot at the end of the iteration. If the left number is ever less than or equal to the pivot, increment the left pointer; similarly, if the right number is ever greater than or equal to the pivot, decrement the right pointer. Do this until the pointers pass each other, at which point swapping the pivot with the right number should position the pivot in its final, sorted position, where every number to its left is smaller and every number to its right is greater. If the pivot is in the kth position, you're done; if it isn't, figure out if the kth smallest number is located to the left or to the right of the pivot.

Hint 3
> Repeat the process mentioned in Hint #2 on the side of the kth smallest number, and keep on repeating the process thereafter until you find the answer. What is the time complexity of this algorithm?

```
Optimal Space & Time Complexity
Best: O(n) time | O(1) space - where n is the length of the input array Average: O(n) time | O(1) space - where n is the length of the input array Worst: O(n^2) time | O(1) space - where n is the length of the input array
```
