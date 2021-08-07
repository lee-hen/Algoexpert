# Largest BST Size

Write a function that takes in a Binary Tree and returns the size of the largest Binary Search Tree (BST) contained in it.

Each BinaryTree node has an integer value, a left child node, and a right child node. Children nodes can either be BinaryTree nodes themselves or None / null.

A BST is a special type of Binary Tree whose nodes all satisfy the BST property. A node satisfies the BST property if its value is strictly greater than the values of every node to its left; its value is less than or equal to the values of every node to its right; and its children nodes are either valid BST nodes themselves or None / null.

## Sample Input

```
tree =            20
          /                \ 
         7                 10
       /   \             /     \
      0     8           5      15
          /   \       /   \   /   \
         7     9     2     5 13   22
                   /           \
                  1             14
```

## Sample Output

```
 9 // The subtree rooted at 10 is the largest BST in the tree, with 9 nodes.
```
