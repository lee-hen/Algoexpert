# Summary

The solution to this question is straightforward from a conceptual point of view but tricky in its implementation.

The overarching idea behind the solution is to traverse through the input binary tree once, to check at every node if the subtree rooted at that node is a BST, and to keep track of the largest found BST.

Traversing through the tree is trivial. As for checking if a subtree is a BST, we need to specifically check that four conditions are met:

   - 1) The left subtree of the current node is itself a BST.
   - 2) The right subtree of the current node is itself a BST.
   - 3) The current node's value is strictly greater than the max value in its left subtree.
   - 4) The current node's value is lesser than or equal to the min value in its right subtree.

This means that, as we traverse the tree, we need each node to return information to its parent node. Specifically, we need each node to provide whether its subtree is a BST, what its subtree's min and max values are, and of course, its subtree's size as well as the size of the largest BST found in its subtree.

Note that the size of each subtree isn't necessarily needed, since we have the running largest BST size at every node, but it arguably makes the code a little easier to read.

## Complexity Analysis

This question is solved in linear time, since we traverse through each node only once and perform only constant-time operations at each node, and with a space complexity of O(h), where h is the height of the tree, which comes from the recursive aspect of the solution.

## Closing Thoughts

This is a classic BST question that assesses your ability to traverse trees, all the while passing information up from children nodes to their parent nodes.
