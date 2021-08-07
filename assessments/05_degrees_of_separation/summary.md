# Summary

The first difficulty that this question introduces is the amount of information that it gives you. In order to grasp what this question is asking, you have to parse a lot of information.

In a real interview, it would be important to ask clarifying questions in order to make sure that you correctly understand what the problem is asking.

In this case, the problem gives us what can essentially be described as a graph of friends as well as two names, and it wants us to return the name (out of the two names) that has the fewest vertices in the graph that are more than six edges away.

This can problem can be solved in a pretty straightforward manner by running two breadth-first searches on the graph, starting at the two people's names, and keeping track of anybody we've already visited in a single BFS as well as the degree (the distance) of anybody we visit.

The number of people who are more than six edges away or who are unvisited in a single BFS run is the number we're looking for; then, we just need to compare this number for the two people we're given.

Note that, since we're doing BFS, we could break out of our traversal as soon as we encounter someone who's seven edges away, since anybody afterwards will naturally be more than six edges away. That being said, while this might speed up our algorithm in practice, it doesn't affect the actual time complexity of our algorithm.

## Complexity Analysis

The complexity analysis for this solution is essentially identical to the complexity analysis of a normal BFS.

## Closing Thoughts

As mentioned above, this question is difficult in large part because of the amount of information that you have to parse in order to get started.

Otherwise, it's just a matter of being comfortable implementing BFS and keeping track of edge distances as well as visited vertices.