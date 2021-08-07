# Balance Index

Write a function that takes in a non-empty array of integers and returns its balance index.

An array a's balance index is an index i such that a[0] + a[1] + ... + a[i - 1] is equal to a[i + 1] + a[i + 2] + ... + a[lastIdx].

In order for 0 to be an array's balance index, a[1] + a[2] + ... + a[lastIdx] must be equal to 0.

Similarly, in order for the last index of an array to be the array's balance index, a[0] + a[1] + ... + a[lastIdx - 1] must be equal to 0.

If there is no balance index in the array, your function should return -1; if there are multiple possible answers, your function can return any of them.

## Sample Input

```
  array = [0, 9, -8, 2, 7, 1, 11, -2, 1]
```

## Sample Output

```
  5 // 0 + 9 + -8 + 2 + 7 == 11 + -2 + 1
```