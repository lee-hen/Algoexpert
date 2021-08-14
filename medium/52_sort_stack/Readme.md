# Sort Stack

Write a function that takes in an array of integers representing a stack, recursively sorts the stack in place (i.e., doesn't create a brand new array), and returns it.

The array must be treated as a stack, with the end of the array as the top of the stack. Therefore, you're only allowed to

   - Pop elements from the top of the stack by removing elements from the end of the array using the built-in .pop() method in your programming language of choice.
   - Push elements to the top of the stack by appending elements to the end of the array using the built-in .append() method in your programming language of choice.
   - Peek at the element on top of the stack by accessing the last element in the array.

You're not allowed to perform any other operations on the input array, including accessing elements (except for the last element), moving elements, etc.. You're also not allowed to use any other data structures, and your solution must be recursive.

## Sample Input

```
stack = [-5, 2, -2, 4, 3, 1]
```

## Sample Output

```
[-5, -2, 1, 2, 3, 4]
```

### Hints

Hint 1
> If you had to insert a single item into an already sorted stack, all the while abiding by the constraints of this problem, how would you do that?

Hint 2
> Inserting a single item in an already sorted stack is fairly simple: you can pop elements off of the stack until you find an element that's smaller than or equal to the value that you want to add. Then, you can push that value on top of the stack and reinsert all the previously popped items back on top of the stack in the reverse order in which you popped them off. The resulting stack will still be sorted.

Hint 3
> You can easily insert multiple items in an already sorted stack by just repeatedly performing what's described in Hint #2. However, you'll need to have an already sorted stack. To get an already sorted stack, you'll need to pop all of the elements off the unsorted stack until it's eventually empty, and then you'll need to push all of the items back on the stack, inserting them in their sorted order one at a time.

Hint 4
> If you're thinking about Hint #3 recursively, the steps are the following:
>
>   1. Pop an item from the top of the stack, and hold onto it in memory.
>   1. Sort the rest of the stack. To do so, repeat step #1 until the stack is empty, at which point you've reached the base case since an empty stack is always sorted.
>   1. Insert the most recently popped off item from step #1 back into the now sorted stack but in its proper sorted position. The first time that you reinsert an item, it'll be inserted in an empty stack.

```
Optimal Space & Time Complexity
O(n^2) time | O(n) space - where n is the length of the stack
```
