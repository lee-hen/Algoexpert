# Glob Matching

In most modern-day computers, glob patterns are used to refer to multiple file names on the computer's system at once.

Glob patterns typically take advantage of the following two special characters:

  - 1) Wildcards, represented by the * symbol, which match any number of characters, including zero characters.
  - 2) Question marks, represented by the ? symbol, which match any single character (exactly one).
   For example, the glob pattern "*.js" matches any file name ending in the JavaScript .js extension.

Write a function that takes in a file name and a pattern (both strings) and returns whether that file name matches the pattern.

## Sample Input
```
fileName = "abcdefg"
pattern = "a*e?g"
```

## Sample Output
```
true
```
