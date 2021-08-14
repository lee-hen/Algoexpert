# Palindrome Partitioning Min Cuts

Given a non-empty string, write a function that returns the minimum number of cuts needed to perform on the string such that each remaining substring is a palindrome.

A palindrome is defined as a string that's written the same forward as backward. Note that single-character strings are palindromes.

## Sample Input

``` 
string = "noonabbad"
```

## Sample Output
``` 
2 // noon | abba | d"
```

### Hints

Hint 1
> Try building a two-dimensional array of the palindromicities of all substrings of the input string. Let the value stored at row i and at column j represent the palindromicity of the substring starting at index i and ending at index j.

Hint 2
> Checking for palindromicity is typically an O(n) time operation. Can you eliminate this step and build the same two-dimensional array mentioned in Hint #1 a different way? Realize that the substring whose starting and ending indices are (i, j) is only a palindrome if string[i] is equal to string[j] and if the substring denoted by (i + 1, j - 1) is also a palindrome.

Hint 3
> Build a one-dimensional array of the same length as the input string. At each index i in this array compute and store the minimum number of cuts needed for the substring whose starting and ending indices are (0, i). Use previously calculated values as well as the two-dimensional array mentioned in Hint #1 to find each value in this array.

```
Optimal Space & Time Complexity
O(n^2) time | O(n^2) space - where n is the length of the input string
```
