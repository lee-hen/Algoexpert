# Reverse Linked List

Write a function that takes in the head of a Singly Linked List, reverses the list in place (i.e., doesn't create a brand new list), and returns its new head.

Each LinkedList node has an integer value as well as a next node pointing to the next node in the list or to None / null if it's the tail of the list.

You can assume that the input Linked List will always have at least one node; in other words, the head will never be None / null.

## Sample Input

```
head = 0 -> 1 -> 2 -> 3 -> 4 -> 5 // the head node with value 0
```

## Sample Output

```
5 -> 4 -> 3 -> 2 -> 1 -> 0 // the new head node with value 5
```
### Hints

```
Hint 1
You can iterate through the Linked List from head to tail and reverse it in place along the way.
```

```
Hint 2
You'll need to manipulate three nodes at once at every step.
```

```
Hint 3
Imagine you have three variables pointing to three consecutive nodes in a Linked List. Start by setting the "next" property of the second node to the first node. Then, set the first variable to the second node, and set the second variable to the third node. Finally, set the third variable to the second variable's "next" property (at this point, the second variable is the original third node). Repeat this process until you're at the tail of the Linked List.
```

```
Optimal Space & Time Complexity
O(n) time | O(1) space - where n is the number of nodes in the Linked List
```
