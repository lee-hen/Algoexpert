# First Duplicate Value

Given an array of integers between 1 and n, inclusive, where n is the length of the array, write a function that returns the first integer that appears more than once (when the array is read from left to right).

In other words, out of all the integers that might occur more than once in the input array, your function should return the one whose first duplicate value has the minimum index.

If no integer appears more than once, your function should return -1.

Note that you're allowed to mutate the input array.

## Sample Input #1

```
array = [2, 1, 5, 2, 3, 3, 4]
```

## Sample Output #1

```
2 // 2 is the first integer that appears more than once.
// 3 also appears more than once, but the second 3 appears after the second 2.
```

## Sample Input #2

```
array = [2, 1, 5, 3, 3, 2, 4]
```

## Sample Output #2
```
3 // 3 is the first integer that appears more than once.
// 2 also appears more than once, but the second 2 appears after the second 3.
```

## Hints

Hint 1
> The brute-force solution can be done in O(n^2) time. Think about how you can determine if a value appears twice in an array.

Hint 2
> You can use a data structure that has constant-time look-ups to keep track of integers that you've seen already. This leads the way to a linear-time solution.

Hint 3
> You should always pay close attention to the details of a question's prompt. In this question, the integers in the array are between 1 and n, inclusive, where n is the length of the input array. The prompt also explicitly allows us to mutate the array. How can these details help us find a better solution, either time-complexity-wise or space-complexity-wise?

Hint 4
> Since the integers are between 1 and the length of the input array, you can map them to indices in the array itself by subtracting 1 from them. Once you've mapped an integer to an index in the array, you can mutate the value in the array at that index and make it negative (by multiplying it by -1). Since the integers normally aren't negative, the first time that you encounter a negative value at the index that an integer maps to, you'll know that you'll have already seen that integer.

```
Optimal Space & Time Complexity
O(n) time | O(1) space - where n is the length of the input array
```
