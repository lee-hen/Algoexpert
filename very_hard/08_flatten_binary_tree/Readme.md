# Flatten Binary Tree

Write a function that takes in a Binary Tree, flattens it, and returns its leftmost node.

A flattened Binary Tree is a structure that's nearly identical to a Doubly Linked List (except that nodes have left and right pointers instead of prev and next pointers), where nodes follow the original tree's left-to-right order.

Note that if the input Binary Tree happens to be a valid Binary Search Tree, the nodes in the flattened tree will be sorted.

The flattening should be done in place, meaning that the original data structure should be mutated (no new structure should be created).

Each BinaryTree node has an integer value, a left child node, and a right child node. Children nodes can either be BinaryTree nodes themselves or None / null.

## Sample Input

``` 
tree =      1
         /     \
        2       3
      /   \   /
     4     5 6
          / \
         7   8
```

## Sample Output

```
4 <-> 2 <-> 7 <-> 5 <-> 8 <-> 1 <-> 6 <-> 3 // the leftmost node with value 4
```

### Hints

Hint 1
> You can solve this problem pretty easily by traversing the tree using the in-order tree-traversal technique, gathering all of the nodes in an array, and then iterating through them from left to right and connecting them accordingly. Can you solve this problem without storing an entire array of the tree's nodes?

Hint 2
> Try to figure out what the relation between two adjacent nodes in the in-order-traversal order is, as far as positioning in the tree is concerned.

Hint 3
> At any given node in the in-order-traversal order, the node immediately to its left is the rightmost node of its left subtree, and the node immediately the its right is the leftmost node of its right subtree.

Hint 4
> Write a function that recursively gets the leftmost and rightmost nodes of a given node's left subtree and right subtree and that connects the left subtree's rightmost node to the given node and the right subtree's leftmost node to the given node.

```
Optimal Space & Time Complexity
O(n) time | O(d) space - where n is the number of nodes in the Binary Tree and d is the depth (height) of the Binary Tree
```
