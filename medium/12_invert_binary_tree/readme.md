# Invert Binary Tree

Write a function that takes in a Binary Tree and inverts it. In other words, the function should swap every left node in the tree for its corresponding right node.

Each BinaryTree node has an integer value, a left child node, and a right child node. Children nodes can either be BinaryTree nodes themselves or None / null.

## Sample Input

```
tree =    1
       /     \
      2       3
    /   \   /   \
   4     5 6     7
 /   \
8     9
```

## Sample Output

```
       1
    /     \
   3       2
 /   \   /   \
7     6 5     4
            /   \
           9     8
```

### Hints

```
Hint 1
Start by inverting the root node of the Binary Tree. Inverting this root node simply consists of swapping its left and right child nodes, which can be done the same way as swapping two variables.
```

```
Hint 2
Once the first swap mentioned in Hint #1 is done, you must invert the root node's left child node and its right child node. You can do so just as you did for the root node. Then, you will have to continue inverting child nodes' nodes until you reach the bottom of the tree.
```

```
Hint 3
How can you accomplish the process described in Hint #2? While recursion seems appropriate, would an iterative approach work? What would be the time and space complexity implications of both approaches?
```

```
Optimal Space & Time Complexity
O(n) time | O(d) space - where n is the number of nodes in the Binary Tree and d is the depth (height) of the Binary Tree
```
