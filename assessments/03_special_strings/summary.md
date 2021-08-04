# Summary

This question can be solved optimally by using a trie.

We start by inserting every string into the trie, and then, for each string, we traverse both the string in question and the trie at the same time to see if at least two instances of other strings are contained in the string in question.

The string and trie traversal is straightforward and can be implemented recursively. At each recursive call on a character in the string, a few things can happen:

   - 1) The character isn't in the current trie node, in which case we break out of the function call.
   - 2) The character is in the current trie node and has an end symbol, meaning that we have a fully formed string; and the character is the last one in the string. In this case, we check that we've counted at least two strings throughout the traversal; if we have, the string is special.
   - 3) The character is in the current trie node and has an end symbol, meaning that we have a fully formed string; and the character isn't the last one in the string. In this case, we recursively check that the rest of the string is special, passing the root of the trie as the next trie node to traverse and incrementing our count of found words in this recursive branch. If the rest of the string isn't special, we proceed to step 4).
   - 4) The character is in the current trie node but doesn't have an end symbol. In this case, we continue our traversal by recursively moving forward in the string and in the trie.
    
We can pretty cleanly write out this algorithm by simply filtering the input list of strings down to the strings that are special, using our recursive function to determine if they are.

## Complexity Analysis

Let n be the number of input strings and m be the length of the longest string.

To insert every string in the trie takes O(n * m) time and O(n * m) space.

Checking if each string is special will also take, on average, O(n * m) time since we'll be traversing through each string roughly once.

The one exception, which is a little tricky to express from a time-complexity point of view, is when we have to repeatedly branch out into two recursive calls of isSpecial. This will only happen for very contrived inputs with many identical prefixes that trigger various recursive-branch explorations; but even then, on average, the time complexity shouldn't exceed O(n * m).

And lastly, our final list of special words will take at most O(n * m) space, bringing the total space-time complexity of the algorithm to O(n * m).

## Closing Thoughts

This question is a perfect example of a string question that can be solved using tries.
