# Min Max Stack Construction

Write a MinMaxStack class for a Min Max Stack. The class should support:

Pushing and popping values on and off the stack.
Peeking at the value at the top of the stack.
Getting both the minimum and the maximum values in the stack at any given point in time.
All class methods, when considered independently, should run in constant time and with constant space.

## Sample Usage
```
// All operations below are performed sequentially.
MinMaxStack(): - // instantiate a MinMaxStack
push(5): -
getMin(): 5
getMax(): 5
peek(): 5
push(7): -
getMin(): 5
getMax(): 7
peek(): 7
push(2): -
getMin(): 2
getMax(): 7
peek(): 2
pop(): 2
pop(): 7
getMin(): 5
getMax(): 5
peek(): 5
```

### Hints

Hint 1
> You should be able to push values on, pop values off, and peek at values on top of the stack at any time and in constant time, using constant space. What data structure maintains order and would allow you to do this?

Hint 2
> You should be able to get the minimum and maximum values in the stack at any time and in constant time, using constant space. What data structure would allow you to do this?

Hint 3
> Since the minimum and maximum values in the stack can change with every push and pop, you will likely need to keep track of all the mins and maxes at every value in the stack.

```
Optimal Space & Time Complexity
All methods: O(1) time | O(1) space
```
