# Balanced Brackets

Write a function that takes in a string made up of brackets ((, [, {, ), ], and }) and other optional characters. The function should return a boolean representing whether the string is balanced with regards to brackets.

A string is said to be balanced if it has as many opening brackets of a certain type as it has closing brackets of that type and if no bracket is unmatched. Note that an opening bracket can't match a corresponding closing bracket that comes before it, and similarly, a closing bracket can't match a corresponding opening bracket that comes after it. Also, brackets can't overlap each other as in [(]).

## Sample Input

```
string = "([])(){}(())()()"
```

## Sample Output
```
true // it's balanced
```

### Hints

Hint 1
> If you iterate through the input string one character at a time, there are two scenarios in which the string will be unbalanced: either you run into a closing bracket with no prior matching opening bracket or you get to the end of the string with some opening brackets that haven't been matched. Can you use an auxiliary data structure to keep track of all the brackets and efficiently check if you run into a unbalanced scenario at every iteration?

Hint 2
> Consider using a stack to store opening brackets as you traverse the string. The Last-In-First-Out property of the stack should allow you to match any closing brackets that you run into against the most recent opening bracket, if one exists, in which case you can simply pop it out of the stack. How can you check that there are no unmatched opening bracket once you've finished traversing through the string?

```
Optimal Space & Time Complexity
O(n) time | O(n) space - where n is the length of the input string
```
