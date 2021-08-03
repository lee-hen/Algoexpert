# Summary

This question is a classic dynamic programming question. If we try to solve it using a brute-force approach, we quickly realize that we'll be doing far too many operations, and even getting a brute-force approach to work is a brain twister.

The most straightforward way to solve the question is to build up a table--a two-dimensional array--containing whether various prefix combinations of the file name and pattern match.

Specifically, the table's dimensions will be (n + 1) * (m + 1), where n and m are the respective lengths of the file name and pattern, and each value in the table at indices (i, j) will represent whether the file name ending at index i - 1 and the pattern ending at index j - 1 match.

The first row and column will represent the empty strings.

The values in the table will be initialized to false except for the very first cell, which represents the two empty strings, and the values in the first row will be incrementally set to true so long the pattern has only * symbols.

Then, as we iterate through the table, row by row and column by column, the value at each cell has three options:

  - 1) Let matchTable [i - 1] [j - 1] be whether the strings at the previous indices match (the previous value); if the pattern at the j - 1th index is a ? or if the pattern at the j - 1th index and the file name at the i - 1th index match, then the value in the current cell is equal to the previous value. In other words, the latest characters in the strings match, so we just check that the strings without the latest characters also match.
  - 2) If the latest pattern character is a * , then the value of the cell is set to whichever value above or to the left of it is true (or false if neither of them are true). For example, if we're at the last indices of abc* and abc, these two strings will match if abc and abc match or if abc* and ab match.
  - 3) If the latest pattern character is neither a ? nor a *, and if the latest two characters aren't equal to each other, then the value is set to false.

The final table for the sample input of this question will look like:

```
      ""      a      *      e      ?      g
""  true, false, false, false, false, false
a  false,  true, true,  false, false, false
b  false, false, true,  false, false, false
c  false, false, true,  false, false, false
d  false, false, true,  false, false, false
e  false, false, true,   true, false, false
f  false, false, true,  false,  true, false
g  false, false, true,  false, false,  true
```

The bottom right corner of the table tells us if the file name and pattern match.

## Complexity Analysis

As for with many dynamic programming questions, we can deduce the time and space complexities of this solution pretty easily since all that we're really doing is creating a two-dimensional array and traversing it, performing constant-time operations at every step.

Since the table's dimensions are based off of the two input strings, the time and space complexities of the solution are both going to be O(n * m), where n and m are the respective lengths of the strings.

## Closing Thoughts

As mentioned above, this is a classic dynamic programming question that is best solved by writing out the table of the two strings and trying to figure out what relations between cells make the most sense.
