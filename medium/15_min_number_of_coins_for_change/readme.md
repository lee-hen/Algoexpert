# Min Number Of Coins For Change

Given an array of positive integers representing coin denominations and a single non-negative integer n representing a target amount of money, write a function that returns the smallest number of coins needed to make change for (to sum up to) that target amount using the given coin denominations.

Note that you have access to an unlimited amount of coins. In other words, if the denominations are [1, 5, 10], you have access to an unlimited amount of 1s, 5s, and 10s.

If it's impossible to make change for the target amount, return -1.

## Sample Input

```
n = 7
denoms = [1, 5, 10]
```

## Sample Output

```
3 // 2x1 + 1x5
```

### Hints

```
Hint 1
Try building an array of the minimum number of coins needed to make change for all amounts between 0 and n inclusive. Note that no coins are needed to make change for 0: in order to make change for 0, you do not need to use any coins.
```

```
Hint 2
Build up the array mentioned in Hint #1 one coin denomination at a time. In other words, find the minimum number of coins needed to make change for all amounts between 0 and n with only one denomination, then with two, etc., until you use all denominations.
```

```
Optimal Space & Time Complexity
O(nd) time | O(n) space - where n is the target amount and d is the number of coin denominations
```