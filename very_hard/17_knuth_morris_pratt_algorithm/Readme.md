# Knuth—Morris—Pratt Algorithm

Write a function that takes in two strings and checks if the first string contains the second one using the Knuth—Morris—Pratt algorithm. The function should return a boolean.

If you're unfamiliar with the Knuth—Morris—Pratt Algorithm, we recommend watching the Conceptual Overview section of this question's video explanation before starting to code.

## Sample Input

```
string = "aefoaefcdaefcdaed"
substring = "aefcdaed"
```

## Sample Output
```
true
```

### Hints

Hint 1
> The Knuth—Morris—Pratt algorithm works by identifying patterns in the potential substring and exploiting them to avoid doing needless character comparisons when searching for the substring in the main string. For instance, take the string "ababac" and the substring "abac"; comparing these strings will fail at the fourth character, where "b" is not equal to "c". Instead of having to restart our comparisons at the second character of the main string, however, we notice that the substring "ab", which is at the beginning of our potential substring, just appeared near our point of failure in the main string. How can we use this to our advantage?

Hint 2
> Start by traversing the potential substring and building out a pattern table. This 1-dimensional array should store, for every position in the substring, the last index at which a matching pattern has been seen; more specifically, this index should be the ending index of a prefix in the substring that is also a suffix at the given position. For example, the string "abcababcd" should yield the following pattern table: [-1, -1, -1, 0, 1, 0, 1, 2, -1].

Hint 3
> After the pattern table mentioned in Hint #2 has been built, traverse the main string and the potential substring with two separate pointers. When characters match, move the pointers forward. When characters don't match, check if the pointer in the substring is at the very beginning of the substring; if it is, then there is no match and you can move the pointer of the main string forward until there is a match; if it isn't, then move it to the position that comes right after the last seen pattern stored at the previous index in the pattern table.

```
Optimal Space & Time Complexity
O(n + m) time | O(m) space - where n is the length of the main string and m is the length of the potential substring
```
