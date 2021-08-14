# Knapsack Problem

You're given an array of arrays where each subarray holds two integer values and represents an item; the first integer is the item's value, and the second integer is the item's weight. You're also given an integer representing the maximum capacity of a knapsack that you have.

Your goal is to fit items in your knapsack without having the sum of their weights exceed the knapsack's capacity, all the while maximizing their combined value. Note that you only have one of each item at your disposal.

Write a function that returns the maximized combined value of the items that you should pick as well as an array of the indices of each item picked.

If there are multiple combinations of items that maximize the total value in the knapsack, your function can return any of them.

## Sample Input

```
items = [[1, 2], [4, 3], [5, 6], [6, 7]]
capacity = 10
```

## Sample Output

```
[10, [1, 3]] // items [4, 3] and [6, 7]
```

### Hints

```
Hint 1
Try building a two-dimensional array of the maximum values that knapsacks of all capacities between 0 and c inclusive could hold, given one, two, three, etc., items. Let columns represent capacities and rows represent items.
```

```
Hint 2
Build up the array mentioned in Hint #1 one row at a time. In other words, find the maximum values that knapsacks of all capacities between 0 and c can hold with only one item, then with two, etc., until you use all items. Find a formula that relates the maximum value at any given point to previous values.
```

```
Hint 3
Backtrack your way through the two-dimensional array mentioned in Hint #1 to find which items are in your knapsack. Start at the final index in the array and check whether or not the value stored at that index is equal to the value located one row above. If it isn't, then the item represented by the current row is in the knapsack.
```

```
Optimal Space & Time Complexity
O(nc) time | O(nc) space - where n is the number of items and c is the capacity
```