# Longest Balanced Substring

Write a function that takes in a string made up of parentheses (( and )). The function should return an integer representing the length of the longest balanced substring with regards to parentheses.

A string is said to be balanced if it has as many opening parentheses as it has closing parentheses and if no parenthesis is unmatched. Note that an opening parenthesis can't match a closing parenthesis that comes before it, and similarly, a closing parenthesis can't match an opening parenthesis that comes after it.

## Sample Input

```
string = "(()))("
```

## Sample Output

```
4 // The longest balanced substring is "(())".
```

### Hints

Hint 1
> With a brute-force style approach, you can iterate through all substrings of the input string, check if they're balanced, and keep track of the longest balanced one. This approach will require using an auxiliary method to check whether a substring is balanced.

Hint 2
> A more efficient approach to solving this problem is to iterate through the input string only once, using a stack to track the indices of all unmatched opening parentheses. Whenever a closing parenthesis is encountered, you check if the stack has a corresponding opening-parenthesis index, and you pop that index off the stack if it does. If the stack doesn't have a corresponding opening-parenthesis index, then the closing parenthesis is unmatched, and its own index in the input string denotes the start of a new, potentially balanced substring. With this approach, you'll have to figure out a way to keep track of how long a balanced substring is.

Hint 3
> The most efficient way to solve this problem is to use only two variables to keep track of the numbers of opening and closing parentheses, respectively, as you traverse the string. Think about how you can use these two pieces of information alone to find the longest balanced substring. Specifically, how do these two pieces of information help you figure out if a substring is balanced, and how can you use them to calculate the length of such a substring?

```
Optimal Space & Time Complexity
O(n) time | O(1) space - where n is the length of the input string
```
