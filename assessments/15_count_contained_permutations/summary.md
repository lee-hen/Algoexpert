# Summary

This question is equivalent to the Smallest Substring Containing question on AlgoExpert; we recommend watching the video explanation of that question to better understand the solution of the question at hand.

There are two main differences here:

   - 1) We're not looking for a single substring in the big string; we're looking for any number of substrings (and more specifically, how many there are).
   - 2) We're not looking for a substring of minimal length; we're looking for substrings of exactly the length of the small string.

In order to solve this problem, we're going to use the same "sliding window" technique that we use to solve the Smallest Substring Containing question, except here our sliding window will constantly have the length of the small string.

This means that as we iterate through the big string with two pointers—a leftIdx and a rightIdx—we're going to make sure that the distance between these two pointers remains equal to the length of the small string; in other words, we're going to move both of the pointers in tandem.

Just as we do in the Smallest Substring Containing question's solution, we're going to constantly keep track of and update the counts of the characters in between our two pointers, and we're going to repeatedly compare these counts to those of the characters in the small string, which we'll precompute.

Whenever the counts of all characters match, we've found a permutation. It's important to note that we can cleverly compare the counts in the constant time by keeping track of the "number of unique characters done"—the number of unique characters in between our two pointers that have counts equal to the relevant ones in the small string. This lets us avoid having to iterate through every unique character in the small string at each iteration in the big string, which would be a linear-time operation with respect to the length of the small string.

The last thing to note is that, in our particular solution, we set up our left and right pointers in the same while loop that we use to move our sliding window. This means that, at first, we only increment our right pointer until it's appropriately distanced from the left pointer. This is done fairly easily with a check that rightIdx - leftIdx == smallString.length.

## Complexity Analysis

Let:

   - b be the length of the big string
   - s be the length of the small string
   
The time complexity of our algorithm is going to be O(b + s), because we're iterating through both strings once, performing constant-time operations at each character.

Note that, since we're told that the small string is smaller than the big string, we could theoretically express this time complexity as just O(b).

The space complexity of our algorithm is going to be O(s), because we're only storing the character counts of the small string as well as those of the substrings between our left and right pointers, which are always separated by a distance of s.

## Closing Thoughts

As mentioned above, this question is a variant of the Smallest Substring Containing question on AlgoExpert. It primarily assesses your ability to use the sliding window technique to efficiently solve questions that involve finding permutations of strings in other strings.

This question also serves as a good test to make sure that a candidate doesn't naively try to compute all permutations of the small string, which would most certainly not lead to an optimal solution.
