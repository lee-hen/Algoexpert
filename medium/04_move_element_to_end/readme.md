# Move Element To End

You're given an array of integers and an integer. Write a function that moves all instances of that integer in the array to the end of the array and returns the array.

The function should perform this in place (i.e., it should mutate the input array) and doesn't need to maintain the order of the other integers.

## Sample Input

```
array = [2, 1, 2, 2, 2, 3, 4, 2]
toMove = 2
```

## Sample Output

```
[1, 3, 4, 2, 2, 2, 2, 2] // the numbers 1, 3, and 4 could be ordered differently
```

### Hints

```
Hint 1
You can solve this problem in linear time.
```

```
Hint 2
In view of Hint #1, you can solve this problem without sorting the input array. Try setting two pointers at the start and end of the array, respectively, and progressively moving them inwards.
```

```
Hint 3
Following Hint #2, set two pointers at the start and end of the array, respectively. Move the right pointer inwards so long as it points to the integer to move, and move the left pointer inwards so long as it doesn't point to the integer to move. When both pointers aren't moving, swap their values in place. Repeat this process until the pointers pass each other.
```

```
Optimal Space & Time Complexity
O(n) time | O(1) space - where n is the length of the array
```
