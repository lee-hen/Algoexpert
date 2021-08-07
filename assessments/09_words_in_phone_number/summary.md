# Summary

This question is equivalent to the Multi String Search question on AlgoExpert and similar to the Boggle Board question; we recommend watching the video explanations of those questions (especially Multi String Search) to better understand this question's solution.

The main difference here is that the big string is a phone number whose digits can represent multiple letters.

We can solve this problem in two ways:

   - 1) By adding all of the words in a trie, iterating through every suffix of the phone number, and for each suffix, traversing through every possible letter combination so long as they're in the trie.
   - 2) By adding all of the phone number's suffixes in a suffix trie, mapping all of the words to their digit equivalents (there's only one word-to-digits possibility per word), and filtering them down to those contained in the suffix tree.
   
The first solution is a bit more complicated than the second one, because we have to go down multiple recursive branches at each digit, and we also have to cleverly avoid doing string concatenations by traversing through the various suffixes with a moving index.

The time complexity of the first solution is also hard to express because it largely depends on how many suffix branches are fully or partially found in the trie.

The second solution, on the other hand, is more straightforward.

## Complexity Analysis
Let:

   - n be the length of the phone number
   - m be the number of words
   - w be the length of the longest word

In the first solution, constructing the word trie takes O(m * w) time and space. Then, there are n suffixes to traverse, and for each suffix, we have roughly three branches at every digit, hence an O(n * 3^n) time operation. As mentioned above, this time complexity is a little tricky to express, because this is the absolute worst case scenario where we have to explore every single possible branch in full, if the trie contains every possible combination of letters. Lastly, the recursive calls of exploreTrie will add O(n) space.

In the second solution, constructing the suffix trie takes O(n^2) time and space, and then, we transform the m words of length at most w into their digit representations--this is done in linear time for each word--and we check if they're contained in the trie--also a linear-time operation. This adds O(m * w) time and O(m * w) space for the final list of words to return.

## Closing Thoughts

As mentioned above, this question is a variant of the Multi String Search question on AlgoExpert and has some similarities to Boggle Board. It primarily assesses your ability to use tries to efficiently solve questions that involve a lot of strings.
