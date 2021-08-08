# Summary

This question can be solved optimally with two different, albeit similar, approaches.

## Approach #1

The general idea behind this approach is to detach relevant groups of k nodes from the main linked list, to reverse these groups using a standard reverseLinkedList method, and to reattach the reversed groups to the main linked list.

Specifically, we traverse through the linked list once, and we keep track of a running counter that gets incremented to k and then gets reset, as well as of a boolean that determines if we're dealing with a group of k nodes that should be reversed (this boolean is initialized to true).

If we reach a node where the running counter is equal to k and the boolean is true, we know that we're dealing with the tail of a group that needs to be reversed, so we detach this tail from the rest of the linked list (by simply setting its next pointer to None / null), and we call our reverseLinkedList method on this group before reattaching it to the main linked list.

Note that we'll have to make sure to keep references to the conceptual tails and heads of the various groups of k nodes in the linked list, since we'll have to update certain pointers when we reattach groups that have been reversed.

Also, we'll have to handle the edge case where the final group of nodes in the linked list to be reversed has fewer than k nodes.

The code in Solution 1 should be self-explanatory.

## Approach #2

This approach is a little simpler to implement than the first approach. The general idea behind this approach is to traverse the linked list, to reverse k nodes with a slightly modified version of the standard reverseLinkedList method, and then to skip the following k nodes by simply iterating through them. We repeat this process until the end of the linked list is reached.

The modified reverseLinkedList method uses a counter to reverse only k nodes, and it returns the head of the reversed linked list (the reversed group of k nodes) as well as the next node in the original linked list.

We still have to make sure to keep references to certain nodes and to attach some nodes correctly, but we don't need to keep track of a running counter or of a boolean determining whether we're at a group that needs to be reversed.

The code in Solution 2 should be self-explanatory.

## Complexity Analysis

This question is self-evidently solved in linear time and with constant space.

## Closing Thoughts

As with most Linked List problems, this question's difficulty lies in handling edge cases, overwriting pointers appropriately, and keeping references to certain nodes.

Oh, and reversing linked lists of course!
