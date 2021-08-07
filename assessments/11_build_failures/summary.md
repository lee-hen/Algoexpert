# Summary

The first difficulty that this question introduces is the amount of information that it gives you. In order to grasp what this question is asking, you have to parse a lot of information.

In a real interview, it would be important to ask clarifying questions in order to make sure that you correctly understand what the problem is asking.

In this case, the problem gives us a list of build runs--lists of booleans specifically structured as [true, true, ..., true, false, false, ..., false]--and it wants us to return the length of the longest stretch of build runs in which the numbers of trues, proportionally speaking, are strictly decreasing.

At face value, it seems like the hard part of this problem is going to be finding this longest stretch.

After all, we can trivially compute the green percentages by finding the index of the first false in each build run and dividing it by the length of the respective build run.

To find the first falses, we can simply iterate through each build run.

...Except, this is actually where we can optimize the solution to this problem. Instead of iterating through each build run, which would be a linear-time operation for each build run, we can cleverly use binary search to find the first falses and to improve this operation to log(n) time.

As far as finding the longest stretch of strictly decreasing green percentages, this is actually a fairly straightforward algorithm to implement; the code should speak for itself.

## Complexity Analysis

We'll have to perform binary search, which is an O(log(m)) operation, where m is the length of the longest build run, n times--once on each build run. So the time complexity of our algorithm is going to be O(nlog(m)).

We can perform the algorithm with constant space by not actually storing every green percentage, but instead only storing the last two green percentages as we iterate through the build runs and find the longest streak.

Implementing our binary search iteratively will ensure that we don't use up auxiliary space because of recursion.

## Closing Thoughts

As mentioned above, this question is difficult in large part because of the amount of information that you have to parse in order to get started.

It's also difficult because it camouflages the step of finding the first false in each build run, which is the meatiest step of the solution--the step where we use binary search. It's very easy to think that the main focus of our algorithm will be on finding the longest streak and to therefore gloss over the optimization with binary search.
