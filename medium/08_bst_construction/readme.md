# BST Construction

Write a BST class for a Binary Search Tree. The class should support:

   - Inserting values with the insert method.
   - Removing values with the remove method; this method should only remove the first instance of a given value.
   - Searching for values with the contains method.

Note that you can't remove values from a single-node tree. In other words, calling the remove method on a single-node tree should simply not do anything.

Each BST node has an integer value, a left child node, and a right child node. A node is said to be a valid BST node if and only if it satisfies the BST property: its value is strictly greater than the values of every node to its left; its value is less than or equal to the values of every node to its right; and its children nodes are either valid BST nodes themselves or None / null.

## Sample Usage
```
// Assume the following BST has already been created:
         10
       /     \
      5      15
    /   \   /   \
   2     5 13   22
 /           \
1            14

// All operations below are performed sequentially.
insert(12):   10
            /     \
           5      15
         /   \   /   \
        2     5 13   22
      /        /  \
     1        12  14

remove(10):   12
            /     \
           5      15
         /   \   /   \
        2     5 13   22
      /           \
     1            14

contains(15): true
```

### Hints

```
Hint 1
As you try to insert, find, or a remove a value into, in, or from a BST, you will have to traverse the tree's nodes. The BST property allows you to eliminate half of the remaining tree at each node that you traverse: if the target value is strictly smaller than a node's value, then it must be (or can only be) located to the left of the node, otherwise it must be (or can only be) to the right of that node.
```

```
Hint 2
Traverse the BST all the while applying the logic described in Hint #1. For insertion, add the target value to the BST once you reach a leaf (None / null) node. For searching, if you reach a leaf node without having found the target value that means the value isn't in the BST. For removal, consider the various cases that you might encounter: the node you need to remove might have two children nodes, one, or none; it might also be the root node; make sure to account for all of these cases.
```

```
Hint 3
What are the advantages and disadvantages of implementing these methods iteratively as opposed to recursively?
```

```
Optimal Space & Time Complexity
Average (all 3 methods): O(log(n)) time | O(1) space - where n is the number of nodes in the BST Worst (all 3 methods): O(n) time | O(1) space - where n is the number of nodes in the BST
```
