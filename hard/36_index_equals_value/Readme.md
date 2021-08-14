# Index Equals Value

Write a function that takes in a sorted array of distinct integers and returns the first index in the array that is equal to the value at that index. In other words, your function should return the minimum index where index == array[index].

If there is no such index, your function should return -1.

## Sample Input

```
array = [-5, -3, 0, 3, 4, 5, 9]
```

## Sample Output

```
3 // 3 == array[3]
```

### Hints

Hint 1
> First think about a simple brute-force approach to solve this problem. What is the time complexity of this approach and what improvements could be made to this time complexity?

Hint 2
> If the brute force solution runs in linear time complexity, then a better solution would have to run in O(log(n)) time. Which algorithm has an O(log(n)) time complexity?

Hint 3
> Implement a variation of binary search to solve this problem. Think about what conditions or checks must be added to search for the desired index-value pair.

Hint 4
> As you perform a variation of binary search on the input array, if the value that you're looking at is smaller than its index, cut the left half of the array from the search space, because all values to the left will be smaller than their corresponding indices; this is guaranteed to be true, since left indices will naturally decrement by 1 each and left values will decrement by at least 1 each due to the array being sorted. Similar logic applies to the right side of the array when the value that you're looking at is greater than its index.

Hint 5
> When you encounter a value that's equal to its index, you'll have to perform some additional logic to make sure that you're not potentially missing other values in the array that are equal to their index and that come before the value that you're looking at.

```
Optimal Space & Time Complexity
O(log(n)) time | O(1) space - where n is the length of the input array
```
