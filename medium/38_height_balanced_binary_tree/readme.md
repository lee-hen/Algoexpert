# Height Balanced Binary Tree

You're given the root node of a Binary Tree. Write a function that returns true if this Binary Tree is height balanced and false if it isn't.

A Binary Tree is height balanced if for each node in the tree, the difference between the height of its left subtree and the height of its right subtree is at most 1.

Each BinaryTree node has an integer value, a left child node, and a right child node. Children nodes can either be BinaryTree nodes themselves or None / null.

## Sample Input

```
tree = 1
     /   \
    2     3
  /   \     \
 4     5     6
     /   \
    7     8
```
## Sample Output

```
true
```

### Hints

Hint 1
> To solve this problem, you'll have to determine if every subtree in the Binary Tree is balanced. Which subtrees do you know will always be balanced?

Hint 2
> To determine if a subtree is balanced, you need to know the height of its left and right subtrees. The only exception to this is if a subtree has no left and right subtrees (i.e., it's just a leaf node); in that case, the subtree must be balanced.

Hint 3
> Recursively calculate the left and right subtree heights from each node. Once you know the heights of a particular node's left and right subtrees, you can determine if the subtree rooted at that node is balanced. If a subtree ever isn't balanced, you can immediately conclude that the entire tree isn't balanced. If you make it through the entire tree without finding any unbalanced subtrees, and if you determine that the heights of the main two subtrees aren't more than 1 apart, then the entire tree is balanced.

```
Optimal Space & Time Complexity
O(n) time | O(h) space - where n is the number of nodes in the binary tree
```
