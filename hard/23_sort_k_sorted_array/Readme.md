# Sort K-Sorted Array

Write a function that takes in a non-negative integer k and a k-sorted array of integers and returns the sorted version of the array. Your function can either sort the array in place or create an entirely new array.

A k-sorted array is a partially sorted array in which all elements are at most k positions away from their sorted position. For example, the array [3, 1, 2, 2] is k-sorted with k = 3, because each element in the array is at most 3 positions away from its sorted position.

Note that you're expected to come up with an algorithm that can sort the k-sorted array faster than in O(nlog(n)) time.

## Sample Input

```
array = [3, 2, 1, 5, 4, 7, 6, 5]
k = 3
```

## Sample Output

```
[1, 2, 3, 4, 5, 5, 6, 7]
```

## Hints

Hint 1
> What does the k parameter tell you? How can you use it to come up with an algorithm that runs in O(nlog(k))?

Hint 2
> Since the input array is k-sorted, try repeatedly sorting k elements at a time and inserting the minimum element of all those k elements into its final sorted position in the array.

Hint 3
> What auxiliary data structure would be helpful to quickly determine the minimum element of k elements?

Hint 4
> As you iterate through the array, use a min-heap to keep track of the most recent k elements. At each iteration, remove the minimum value from the heap, insert it into its final sorted position in the array, and add the current element in the array to the heap. Continue this process until the heap is empty.

```
Optimal Space & Time Complexity
O(nlog(k)) time | O(k) space - where n is the number of elements in the array and k is how far away elements are from their sorted position
```
