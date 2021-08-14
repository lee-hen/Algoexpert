# Powerset

Write a function that takes in an array of unique integers and returns its powerset.

The powerset P(X) of a set X is the set of all subsets of X. For example, the powerset of [1,2] is [[], [1], [2], [1,2]].

Note that the sets in the powerset do not need to be in any particular order.

## Sample Input

```
array = [1, 2, 3]
```

## Sample Output

```
[[], [1], [2], [3], [1, 2], [1, 3], [2, 3], [1, 2, 3]]
```

### Hints

```
Hint 1
Try thinking about the base cases. What is the powerset of the empty set? What is the powerset of sets of length 1?
```

```
Hint 2
If you were to take the input set X and add an element to it, how would the resulting powerset change?
```

```
Hint 3
Can you solve this problem recursively? Can you solve it iteratively? What are the advantages and disadvantages of using either approach?
```

```
Optimal Space & Time Complexity
O(n*2^n) time | O(n*2^n) space - where n is the length of the input array
```
