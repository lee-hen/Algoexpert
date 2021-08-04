# Special Strings

Write a function that takes in a list of non-empty strings and returns a list of all the special strings found in the input list.

A string is said to be special if it's exactly made up of at least two instances of other strings in the input list of strings.

In order for a string to be special, just containing two instances of other strings isn't sufficient; the string must be exactly made up of those other strings. For example, in the list ["a", "b", "abc"], the string "abc" isn't special, even though it contains "a" and "b", because "c" isn't a string in the list.

Note that strings can be repeated to form a special string; for instance, in the list ["a", "aaa"], the string "aaa" is a special string because it's made up of three repeated instances of "a".

Also note that you can't use language-built-in string-matching methods.

## Sample Input
``` 
strings = [
  "foobarbaz",
  "foo",
  "bar",
  "foobarfoo",
  "baz",
  "foobaz",
  "foofoofoo",
  "foobazar",
]
```

## Sample Output
```
["foobarbaz", "foobarfoo", "foobaz", "foofoofoo"]
```
