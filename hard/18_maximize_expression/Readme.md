# Maximize Expression

Write a function that takes in an array of integers and returns the largest possible value for the expression array[a] - array[b] + array[c] - array[d], where a, b, c, and d are indices of the array and a < b < c < d.

If the input array has fewer than 4 elements, your function should return 0.

## Sample Input

```
array = [3, 6, 1, -3, 2, 7]
```

## Sample Output

```
4
// Choose a = 1, b = 3, c = 4, and d = 5
// -> 6 - (-3) + 2 - 7 = 4
```

### Hints

Hint 1
> The brute-force approach to solving this problem is to simply iterate through every valid choice of a, b, c, and d and to evaluate the expression at each iteration. While doing this, you can keep track of the maximum value that you find and return it after considering all possibilities. This solution runs in O(n^4) time; can you think of a way to solve this faster?

Hint 2
> You can solve this problem using dynamic programming with a time complexity of O(n); however, you'll need to use external space.

Hint 3
> If you know what the maximum possible value of a is at each index in the array, you can find the maximum possible value of a - b at each individual index in the array in O(1) time (or in O(n) time for all indices). The same thing holds for finding the maximum possible value of a - b + c if you know the maximum possible value of a - b at each index. How does this fact help you solve the entire problem in O(n) time?

Hint 4
> Start by finding the maximum possible value of a at each index in the array, meaning the maximum value of a that you can obtain at each index i if a is chosen from an index between 0 and i, inclusive. Store all of these values in an array, and use them to help you determine the maximum possible value of a - b at each index. Do the same for a - b + c (using the results from a - b) and a - b + c - d (using the results from a - b + c). Once you make it to a - b + c - d, you'll be able to determine the maximum value of the expression.

```
Optimal Space & Time Complexity
O(n) time | O(n) space - where n is the length of the array
```
