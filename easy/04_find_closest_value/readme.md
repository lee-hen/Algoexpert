# Find Closest Value In BST

Write a function that takes in a Binary Search Tree (BST) and a target integer value and returns the closest value to that target value contained in the BST.

You can assume that there will only be one closest value.

Each BST node has an integer value, a left child node, and a right child node. A node is said to be a valid BST node if and only if it satisfies the BST property: its value is strictly greater than the values of every node to its left; its value is less than or equal to the values of every node to its right; and its children nodes are either valid BST nodes themselves or None / null.

## Sample Input

```
   tree =    10
           /     \
          5      15
        /   \   /   \
       2     5 13   22
     /           \
    1            14
    
   target = 12
```

## Sample Output

```
13
```

### Hints

```
Hint 1
Try traversing the BST node by node, all the while keeping track of the node with the value closest to the target value. Calculating the absolute value of the difference between a node's value and the target value should allow you to check if that node is closer than the current closest one.
```

```
Hint 2
Make use of the BST property to determine what side of any given node has values close to the target value and is therefore worth exploring.
```

```
Hint 3
What are the advantages and disadvantages of solving this problem iteratively as opposed to recursively?
```

```
Optimal Space & Time Complexity
Average: O(log(n)) time | O(1) space - where n is the number of nodes in the BST Worst: O(n) time | O(1) space - where n is the number of nodes in the BST
```
