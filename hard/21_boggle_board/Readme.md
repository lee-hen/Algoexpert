# Boggle Board

You're given a two-dimensional array (a matrix) of potentially unequal height and width containing letters; this matrix represents a boggle board. You're also given a list of words.

Write a function that returns an array of all the words contained in the boggle board. The final words don't need to be in any particular order.

A word is constructed in the boggle board by connecting adjacent (horizontally, vertically, or diagonally) letters, without using any single letter at a given position more than once; while a word can of course have repeated letters, those repeated letters must come from different positions in the boggle board in order for the word to be contained in the board. Note that two or more words are allowed to overlap and use the same letters in the boggle board.

## Sample Input
```
board = [
    ["t", "h", "i", "s", "i", "s", "a"],
    ["s", "i", "m", "p", "l", "e", "x"],
    ["b", "x", "x", "x", "x", "e", "b"],
    ["x", "o", "g", "g", "l", "x", "o"],
    ["x", "x", "x", "D", "T", "r", "a"],
    ["R", "E", "P", "E", "A", "d", "x"],
    ["x", "x", "x", "x", "x", "x", "x"],
    ["N", "O", "T", "R", "E", "-", "P"],
    ["x", "x", "D", "E", "T", "A", "E"],
],
words = [
    "this", "is", "not", "a", "simple", "boggle",
    "board", "test", "REPEATED", "NOTRE-PEATED",
]
```

## Sample Output
```
["this", "is", "a", "simple", "boggle", "board", "NOTRE-PEATED"]
// The words could be ordered differently.
```

### Hints

Hint 1
> You can divide this question into two separate problems: one part involves traversing the boggle board in such a way that allows you to construct strings letter by letter; the other part involves actually comparing the strings you construct in the board against the words in the list that you're given. For the second part, what data structure lends itself very well to matching characters to multiple strings at once?

Hint 2
> Try creating a trie out of the input list of words. This will allow you to compare letters in the boggle board against all input words in constant time. How can you efficiently traverse the boggle board to construct all potentially valid strings, without counting letters twice in any string?

Hint 3
> Treat the board as a graph, where each element in the board is a node with up to 8 neighboring nodes. Traverse it in a depth-first-search-like fashion, checking if letters are contained in the trie and traversing the trie simultaneously if it makes sense to do so. How can you keep track of letters that you've already visited in order to avoid erroneously counting some of them twice in a single string? Could you keep track of visited nodes in an auxiliary data structure?

Hint 4
> Keeping in mind that you only want to mark nodes as visited in a single branch of the graph that you're traversing (i.e., you don't want the state of visited nodes in one branch of the graph to spill into the state of another branch of the graph), try marking any node you traverse as unvisited at the end of the recursive call that actually traverses it, after traversing through all of the node's neighbors and performing the same actions on them recursively.

```
Optimal Space & Time Complexity
O(nm*8^s + ws) time | O(nm + ws) space - where n is the width the board, m is the height of the board, w is the number of words, and s is the length of the longest word
```