# Longest Common Subsequence

Write a function that takes in two strings and returns their longest common subsequence.

A subsequence of a string is a set of characters that aren't necessarily adjacent in the string but that are in the same order as they appear in the string. For instance, the characters ["a", "c", "d"] form a subsequence of the string "abcd", and so do the characters ["b", "d"]. Note that a single character in a string and the string itself are both valid subsequences of the string.

You can assume that there will only be one longest common subsequence.

## Sample Input

```
str1 = "ZXVVYZW"
str2 = "XKYKZPW"
```

## Sample Output

```
["X", "Y", "Z", "W"]
```

### Hints

Hint 1
> Try building a two-dimensional array of the longest common subsequences of substring pairs of the input strings. Let the rows of the array represent substrings of the second input string str2. Let the first row represent the empty string. Let each row i thereafter represent the substrings of str2 from 0 to i, with i excluded. Let the columns similarly represent the first input string str1.

Hint 2
> Build up the array mentioned in Hint #1 one row at a time. In other words, find the longest common subsequences for all the substrings of str1 represented by the columns and the empty string represented by the first row, then for all the substrings of str1 represented by the columns and the first letter of str2 represented by the second row, etc., until you compare both full strings. Find a formula that relates the longest common subsequence at any given point to previous subsequences.

Hint 3
> Do you really need to build and store subsequences at each point in the two-dimensional array mentioned in Hint #1? Try storing booleans to determine whether or not a letter at a given point in the two-dimensional array is part of the longest common subsequence as well as pointers to determine what should come before this letter in the final subsequence. Use these pointers to backtrack your way through the array and to build up the longest common subsequence at the end of your algorithm.

```
Optimal Space & Time Complexity
O(nm) time | O(nm) space - where n and m are the lengths of the two input strings
```
